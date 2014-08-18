package models

type Model interface {
	Construct(arg ...interface{}) interface{}
}
