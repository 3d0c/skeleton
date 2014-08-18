package utils

var AppConfig *ConfigScheme

func init() {
	AppConfig = &ConfigScheme{}
}

type ConfigScheme struct {
	App struct {
		ListenOn string `json:"listen_on"`
		HttpsOn  string `json:"https_on"`
		SSLCert  string `json:"ssl_cert"`
		SSLKey   string `json:"ssl_key"`
	} `json:"application"`
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
