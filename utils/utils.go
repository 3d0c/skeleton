package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

func HashOf(in string) string {
	secretKey := []byte("CHANGE_ME!")

	mac := hmac.New(sha1.New, secretKey)

	io.WriteString(mac, in)

	return hex.EncodeToString(mac.Sum(nil))
}
