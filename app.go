package main

import (
	"fmt"
	ctrl "github.com/3d0c/skeleton/controllers"
	"github.com/3d0c/skeleton/models"
	. "github.com/3d0c/skeleton/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	if err := InitConfigFrom("./config.json"); err != nil {
		log.Fatalln("Unable to proceed")
	}
}

func main() {
	m := martini.New()
	route := martini.NewRouter()

	// Some CORS stuff
	m.Use(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			w.Header().Add("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Add("Access-Control-Allow-Origin", "*")
		}

		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Add("Cache-Control", "max-age=2592000")
		w.Header().Add("Pragma", "public")
		w.Header().Add("Cache-Control", "public")
	})

	m.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))

	// Preflight OPTIONS
	route.Options("/**")

	route.Get("/user",
		construct(&models.User{}),
		construct(&ctrl.User{}),
		(*ctrl.User).Find,
	)

	route.Get("/posts",
		construct(&models.User{}),
		construct(&ctrl.Post{}),
		(*ctrl.Post).FindAll,
	)

	route.Get("/posts/:id",
		construct(&models.User{}),
		construct(&ctrl.Post{}),
		(*ctrl.Post).Find,
	)

	m.Action(route.Handle)

	log.Printf("Waiting for connections on %v...\n", AppConfig.ListenOn())

	go func() {
		if err := http.ListenAndServe(AppConfig.ListenOn(), m); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServeTLS(AppConfig.HttpsOn(), AppConfig.SSLCert(), AppConfig.SSLKey(), m); err != nil {
		log.Fatal(err)
	}
}

func construct(obj interface{}, args ...interface{}) martini.Handler {
	return func(ctx martini.Context, r *http.Request) {
		switch t := obj.(type) {
		case models.Model:
			ctx.Map(obj.(models.Model).Construct(args))

		case ctrl.Controller:
			ctx.Map(obj.(ctrl.Controller).Construct(args))

		default:
			panic(fmt.Sprintln("Unexpected type:", t))
		}
	}
}
