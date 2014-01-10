package controllers

import (
	"github.com/3d0c/skeleton/models"
	"net/http"
)

func CheckAuth(u models.Users, res http.ResponseWriter) {
	if !u.Authenticated() {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte{})
	}
}
