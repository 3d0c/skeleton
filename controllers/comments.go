package controllers

import (
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/models"
	// "github.com/codegangsta/martini"
	// "labix.org/v2/mgo/bson"
	"net/http"
)

func CommentsCreate(u models.Users, post models.Posts, comment models.Comments, params models.CommentScheme, enc encoder.Encoder) (int, []byte) {
	var result interface{}

	// we've got here a post model, so we can check
	// is there a post with 'id', given by user.

	params.Uid = u.Object.Id

	if result = comment.Create(&params); result == nil {
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}
