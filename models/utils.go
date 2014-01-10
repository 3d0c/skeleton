package models

import (
	"github.com/codegangsta/martini"
	"log"
	"net/http"
	"reflect"
)

// Creates instance of {obj} and invokes its {Constructor} method with given args.
func Construct(obj interface{}, args ...interface{}) martini.Handler {
	return func(context martini.Context, req *http.Request) {
		obj := reflect.New(reflect.TypeOf(obj))

		if method := obj.MethodByName("Construct"); method.IsValid() {
			instance := method.Call([]reflect.Value{reflect.ValueOf(args)})[0]
			if instance.Kind() == reflect.Ptr {
				context.Map(instance.Elem().Interface())
			} else {
				context.Map(instance.Interface())
			}
		} else {
			log.Println("Construct method not found or invalid")
		}
	}
}
