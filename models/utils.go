package models

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/codegangsta/martini"
	"io"
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

func HashOf(in string) string {
	secretKey := []byte("CHANGE_ME!")

	mac := hmac.New(sha1.New, secretKey)

	io.WriteString(mac, in)

	return hex.EncodeToString(mac.Sum(nil))
}
