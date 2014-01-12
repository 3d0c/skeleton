package models

import (
	"github.com/3d0c/martini-contrib/model"
)

type Comments struct {
	*model.Model
}

func (*Comments) Construct(args ...interface{}) *Comments {
	this := &Comments{
		Model: model.New(CommentScheme{}),
	}

	return this
}
