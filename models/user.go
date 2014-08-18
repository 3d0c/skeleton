package models

import (
	"log"
)

type UserScheme struct {
	Name string `json:"name"`
}

type User struct {
	Object *UserScheme
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{
		Object: &UserScheme{Name: "test one"},
	}

	log.Println("User model initialized,", this)

	return this
}

func (this *User) Name() string {
	return this.Object.Name
}
