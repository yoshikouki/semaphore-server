package server

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port int `env:"PORT" envDefault:"8686"`
}

func NewConfig() (Config, error) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		fmt.Errorf("NewConfig is Error: %+v\n", err)
	}
	return conf, nil
}
