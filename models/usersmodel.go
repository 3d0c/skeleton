package models

import (
	"github.com/3d0c/martini-contrib/model"
	"github.com/3d0c/skeleton/utils"
	"log"
)

type Users struct {
	*model.Model
	Object *UserScheme
}

func (*Users) Construct(args ...interface{}) *Users {
	this := &Users{
		Model:  model.New(UserScheme{}),
		Object: &UserScheme{},
	}

	v := args[0].([]interface{})

	for _, value := range v {
		switch t := value.(type) {
		case *utils.Credentials:
			if value.(*utils.Credentials).Got() {
				this.Object = this.Find(this.authQuery(value.(*utils.Credentials))).One().(*UserScheme)
			}
			break

		default:
			log.Println("Unknown type:", t)
			break
		}
	}

	return this
}

func (this *Users) Authenticated() bool {
	if this.Object.Id == "" {
		return false
	}

	return true
}

func (this *Users) Enabled() bool {
	if this.Object.Id != "" && this.Object.Status == "enabled" {
		return true
	}

	return false
}

func (this *Users) CustomMethod() {
	log.Println("do something")
}

func (this *Users) authQuery(credentials *utils.Credentials) interface{} {
	return map[string]interface{}{
		"login":    credentials.User(),
		"password": HashOf(credentials.Password()),
	}
}
