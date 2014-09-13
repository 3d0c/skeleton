package controllers

import (
	"github.com/3d0c/skeleton/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

type Post struct {
	model *models.Post
}

func (*Post) Construct(args ...interface{}) interface{} {
	this := &Post{
		model: (*models.Post).Construct(nil).(*models.Post),
	}

	log.Println("Post controller initialized,", this)

	return this
}

func (this *Post) Find(u *models.User, r render.Render, p martini.Params) {
	r.JSON(http.StatusOK, this.model.Find(p["id"]))
}

func (this *Post) FindAll(u *models.User, r render.Render) {
	r.JSON(http.StatusOK, this.model.FindAll())
}
