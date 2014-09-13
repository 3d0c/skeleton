package controllers

import (
	"github.com/3d0c/skeleton/models"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
)

type User struct {
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{}
	log.Println("User controller initialized,", this)

	return this
}

func (this *User) Find(u *models.User, r render.Render) {
	r.JSON(http.StatusOK, u.Object)
}
