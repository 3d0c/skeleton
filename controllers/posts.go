package controllers

import (
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/models"
	"github.com/codegangsta/martini"
	"labix.org/v2/mgo/bson"
	// "log"
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

func PostsFind(posts models.Posts, users models.Users, enc encoder.Encoder, urlParams martini.Params, opt models.UrlOptions) (int, []byte) {
	var result interface{}

	if id, ok := urlParams["id"]; ok && bson.IsObjectIdHex(id) {
		result = posts.Find(id).One()
	} else {
		result = posts.Find().Skip(opt.Offset).Limit(opt.Limit).All()
	}

	if result == nil {
		return http.StatusNotFound, []byte{}
	}

	if opt.Expand != "" {
		users.Expand(result, opt.Expand)
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func PostsUpdate(u models.Users, posts models.Posts, enc encoder.Encoder, params models.PostScheme, urlParams martini.Params) (int, []byte) {
	var result interface{}
	selector := map[string]interface{}{}

	if id, ok := urlParams["id"]; ok && bson.IsObjectIdHex(id) {
		selector["_id"] = bson.ObjectIdHex(id)
	} else {
		return http.StatusBadRequest, []byte{}
	}

	selector["uid"] = u.Object.Id

	if result, _ = posts.Update(selector, params); result == nil {
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func PostsDelete(u models.Users, posts models.Posts, urlParams martini.Params) (int, []byte) {
	selector := map[string]interface{}{}

	if id, ok := urlParams["id"]; ok && bson.IsObjectIdHex(id) {
		selector["_id"] = bson.ObjectIdHex(id)
	} else {
		return http.StatusBadRequest, []byte{}
	}

	selector["uid"] = u.Object.Id

	if !posts.Delete(selector) {
		return http.StatusBadRequest, []byte{}
	}

	return http.StatusOK, []byte{}
}
