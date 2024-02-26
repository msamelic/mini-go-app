package envconf

import (
	"github.com/kelseyhightower/envconfig"
)

type Spec struct {
	AppName string `envconfig:"APP_NAME" default:"mini-go-app"`
	Port    string `envconfig:"PORT" default:"8080"`
}

func New() *Spec {
	var s Spec
	envconfig.MustProcess("", &s)
	return &s
}
