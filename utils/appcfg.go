package utils

var AppCfg *ConfigScheme

func init() {
	AppCfg = &ConfigScheme{}
}

type ConfigScheme struct {
	App struct {
		Secretkey string `json:"secretkey"`
		ListenOn  string `json:"listen_on"`
	} `json:"application"`
}

func (this *ConfigScheme) SecretKey() string {
	return this.App.Secretkey
}

func (this *ConfigScheme) ListenOn() string {
	return this.App.ListenOn
}
