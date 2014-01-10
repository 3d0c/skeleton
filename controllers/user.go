package controllers

import (
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/models"
	"github.com/3d0c/skeleton/utils"
	"net/http"
)

func UsersCreate(u models.Users, params models.UserScheme, enc encoder.Encoder) (int, []byte) {
	var result interface{}

	// all params have been already validated
	// here we do some preparation, e.g.:
	params.Password = utils.HashOf(params.Password)

	if result = u.Create(&params); result == nil {
		// we do not disclose internal errors, just return "500 Server Error" and empty body
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

func UserFind(u models.Users, enc encoder.Encoder) (int, []byte) {
	// user object is already loaded by Construct call
	// so we just return it
	return http.StatusOK, encoder.Must(enc.Encode(u.Object))
}

func UserUpdate(u models.Users, params models.UserScheme, enc encoder.Encoder) (int, []byte) {
	var result interface{}

	params.Login = "" // We should prevent user to modify it. (Empty field will be omitted by "omitempty" rule)

	if params.Password != "" {
		params.Password = utils.HashOf(params.Password)
	}

	if result, _ = u.Update(u.Object.Id, &params); result == nil {
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}
