package controllers

type Controller interface {
	Construct(arg ...interface{}) interface{}
}
