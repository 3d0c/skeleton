package main

import (
	"github.com/3d0c/martini-contrib/binding"
	"github.com/3d0c/martini-contrib/config"
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/controllers"
	"github.com/3d0c/skeleton/models"
	"github.com/3d0c/skeleton/utils"
	"github.com/codegangsta/martini"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	config.Init("./app.json")

	m := martini.New()
	route := martini.NewRouter()

	credentials := &utils.Credentials{}
	m.Use(credentials.Get)

	// map json encoder
	m.Use(func(c martini.Context, w http.ResponseWriter) {
		c.MapTo(encoder.JsonEncoder{}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	})

	// some CORS stuff
	m.Use(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origin != "" {
			w.Header().Add("Access-Control-Allow-Origin", origin)
		}

		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
	})

	route.Get("/user",
		models.Construct(models.Users{}, credentials), // Initialize Users model with credentials. It will be available inside controller.
		controllers.CheckAuth,                         // just a helper. controllers/common.go
		controllers.UserFind,                          // func(u models.Users, enc encoder.Encoder)
	)

	route.Put("/user",
		models.Construct(models.Users{}, credentials),
		binding.Bind(models.UserScheme{}),
		controllers.CheckAuth,
		controllers.UserUpdate,
	)

	route.Post("/users",
		binding.Bind(models.UserScheme{}), // Bind parameters
		models.Construct(models.Users{}),  // Init the model without credentials.
		controllers.UsersCreate,           // func(u models.Users, params models.UserScheme, enc encoder.Encoder)
	)

	// Find Id
	route.Get("/posts/:id",
		models.Construct(models.Posts{}),     // Public method. It dosn't need authorization. Init only Posts model.
		binding.BindUrl(models.UrlOptions{}), // For further purposes (expand, etc...)
		controllers.PostsFind,                //
	)

	// Find all
	route.Get("/posts",
		models.Construct(models.Posts{}),     // Public method. It dosn't need authorization. Init only Posts model.
		binding.BindUrl(models.UrlOptions{}), // Bind url options, e.g.: ?limit=10&offset=100, etc... Just a hack,
		controllers.PostsFind,                // actually it invokes patched Form midleware directly (without content-type checking).
	)

	route.Post("/posts",
		models.Construct(models.Users{}, credentials), // Init Users model, because
		models.Construct(models.Posts{}),              // each Post should containt user id.
		binding.Bind(models.PostScheme{}),
		controllers.CheckAuth,
		controllers.PostsCreate,
	)

	route.Put("/posts/:id",
		models.Construct(models.Users{}, credentials), // Init Users model, because
		models.Construct(models.Posts{}),              // each Post should containt user id.
		binding.Bind(models.PostScheme{}),
		controllers.CheckAuth,
		controllers.PostsUpdate,
	)

	route.Delete("/posts/:id",
		models.Construct(models.Users{}, credentials),
		models.Construct(models.Posts{}),
		controllers.CheckAuth,
		controllers.PostsDelete,
	)

	route.Post("/comments",
		models.Construct(models.Users{}, credentials),
		models.Construct(models.Posts{}),
		models.Construct(models.Comments{}),
		binding.Bind(models.CommentScheme{}),
		controllers.CheckAuth,
		controllers.CommentsCreate,
	)

	m.Action(route.Handle)

	log.Println("Waiting for connections...")

	if err := http.ListenAndServe(":8000", m); err != nil {
		log.Fatal(err)
	}
}
