package models

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

type Credentials struct {
	payload interface{}
}

func (this *Credentials) Init(req *http.Request) *Credentials {
	s := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		log.Println("No Authorization header found.")
		return this
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		log.Println(err)
		return this
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		log.Println("Wrong authorization credentials.")
		return this
	}

	// return pair[0], pair[1], nil

	this.payload = []string{pair[0], pair[1]}

	return this
}

func (this *Credentials) GetPair() []string {
	return []string{}
}

func (this *Credentials) GetToken() string {
	return ""
}
