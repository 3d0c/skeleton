package models

import (
	"github.com/codegangsta/martini-contrib/binding"
	"labix.org/v2/mgo/bson"
	"net/http"
)

// Create unique index on Users collection. db.Users.ensureIndex({'login':1},{uniq:true})
//
// Some additional tags:
// 		if a tag "out" sets to "false" this field won't be encoded by encoder package (see copyStruct function).
// 		if a tag "binding" sets to "-", this field won't be binded and will be omitted.
//
type UserScheme struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"        binding:"-"`
	Login    string        `bson:",omitempty"    json:"login"`
	Password string        `bson:",omitempty"    json:"password,omitempty"  out:"false"`
	Status   string        `bson:",omitempty"    json:"status,omitempty"    binding:"-"` // We don't want, that user changed its status.
	Profile  struct {
		FirstName   string `bson:",omitempty"    json:"first_name"`
		LastName    string `bson:",omitempty"    json:"last_name"`
		HiddenField string `bson:",omitempty"    json:"hidden_field,omitempty" out:"false"`
	} `bson:",omitempty" json:"profile"`
}

func (this UserScheme) Validate(errors *binding.Errors, req *http.Request) {
	if len(this.Login) < 4 {
		errors.Fields["login"] = "Too short; minimum 4 characters"
	}
}

type PostScheme struct {
	Id    bson.ObjectId `bson:"_id,omitempty" json:"id" binding:"-"`
	Title string        `bson:",omitempty" json:"title"`
	Body  string        `bson:",omitempty" json:"body"`
	Uid   bson.ObjectId `bson:",omitempty" json:"uid" binding:"-"`
}

func (this PostScheme) Validate(errors *binding.Errors, req *http.Request) {
	if len(this.Title) == 0 {
		errors.Fields["title"] = "Title can't be empty."
	}
}

// This isn't a real model scheme, but there is a good place to store it all together
type UrlOptions struct {
	Limit  int `url:"limit"`
	Offset int `url:"offset"`
}
