package utils

// In opposite to martini-contrib/auth this stuff doesn't check anything,
// it just returns parsed credentials, to process it further inside models.

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

type Credentials struct {
	pair []string
}

func (this *Credentials) Get(req *http.Request) *Credentials {
	s := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 || s[0] != "Basic" {
		return this
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		log.Println(err)
		return this
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return this
	}

	this.pair = []string{pair[0], pair[1]}

	return this
}

func (this *Credentials) Got() bool {
	return len(this.pair) == 2
}

func (this *Credentials) User() string {
	return this.pair[0]
}

func (this *Credentials) Password() string {
	return this.pair[1]
}
