package controllers

import (
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/models"
	"github.com/codegangsta/martini"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
)

func PostsCreate(u models.Users, post models.Posts, params models.PostScheme, enc encoder.Encoder) (int, []byte) {
	var result interface{}

	params.Uid = u.Object.Id

	if result = post.Create(&params); result == nil {
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func PostsFindId(posts models.Posts, enc encoder.Encoder, params martini.Params, req *http.Request) (int, []byte) {
	var result interface{}

	if !bson.IsObjectIdHex(params["id"]) {
		return http.StatusBadRequest, []byte{}
	}

	if result = posts.Find(bson.ObjectIdHex(params["id"])).One(); result == nil {
		return http.StatusNotFound, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func PostsFindAll(posts models.Posts, enc encoder.Encoder, urlParams martini.Params, req *http.Request) (int, []byte) {
	var result interface{}

	if result = posts.Find().All(); result == nil {
		return http.StatusNotFound, []byte{}
	}

	log.Println(result)

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func PostsFind(u models.Users, p models.Posts, enc encoder.Encoder, urlParams martini.Params, req *http.Request) (int, []byte) {
	log.Println(urlParams)
	log.Println(req.URL.Query().Get("limit"))

	result := p.Find()

	log.Println(result)

	return http.StatusOK, encoder.Must(enc.Encode(result))

	// return http.StatusOK, []byte{}
}
