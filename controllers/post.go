package controllers

import (
	"github.com/3d0c/skeleton/models"
	"github.com/martini-contrib/encoder"
	"log"
	"net/http"
)

type Post struct {
}

func (*Post) Construct(args ...interface{}) interface{} {
	this := &Post{}

	log.Println("Post controller initialized,", this)

	return this
}

func (this *Post) Find(u *models.User, enc encoder.Encoder) (int, []byte) {
	return http.StatusOK, []byte{}
}
