package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var AppConfig *ConfigScheme

type ConfigScheme struct {
	App struct {
		ListenOn string `json:"listen_on"`
		HttpsOn  string `json:"https_on"`
		SSLCert  string `json:"ssl_cert"`
		SSLKey   string `json:"ssl_key"`
	} `json:"application"`
}

func InitConfigFrom(file string) error {
	if AppConfig != nil {
		return nil
	}

	AppConfig = &ConfigScheme{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("Unable to read", file, "Error:", err)
		return err
	}

	err = json.Unmarshal(data, AppConfig)
	if err != nil {
		log.Println("Unable to read config.", err)
		return err
	}

	return nil
}

func (this *ConfigScheme) ListenOn() string {
	return this.App.ListenOn
}

func (this *ConfigScheme) HttpsOn() string {
	return this.App.HttpsOn
}

func (this *ConfigScheme) SSLCert() string {
	return this.App.SSLCert
}

func (this *ConfigScheme) SSLKey() string {
	return this.App.SSLKey
}
