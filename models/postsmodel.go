package models

import (
	"github.com/3d0c/martini-contrib/model"
)

type Posts struct {
	*model.Model
}

func (*Posts) Construct(args ...interface{}) *Posts {
	this := &Posts{
		Model: model.New(PostScheme{}),
	}

	return this
}

func (this *Posts) Custom() {
	// do something
}
