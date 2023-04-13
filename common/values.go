package common

import (
	"fmt"
	"github.com/weplanx/utils/values"
	"strings"
)

type Values struct {
	Server                string `env:"SERVER" envDefault:"admin"`
	Address               string `env:"ADDRESS" envDefault:":3000"`
	Namespace             string `env:"NAMESPACE,required"`
	Key                   string `env:"KEY,required"`
	Admin                 `envPrefix:"ADMIN_"`
	Database              `envPrefix:"DATABASE_"`
	Nats                  `envPrefix:"NATS_"`
	*values.DynamicValues `env:"-"`
}

type Admin struct {
	Address string `env:"ADDRESS" envDefault:":3001"`
	Url     string `env:"URL,required"`
}

type Database struct {
	Host  string `env:"HOST,required"`
	Name  string `env:"NAME,required"`
	Redis string `env:"REDIS,required"`
}

type Nats struct {
	Hosts []string `env:"HOSTS,required" envSeparator:","`
	Nkey  string   `env:"NKEY,required"`
}

func (x Values) Name(v ...string) string {
	return fmt.Sprintf(`%s:%s`, x.Namespace, strings.Join(v, ":"))
}