package server

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port          int    `env:"PORT" envDefault:"8686"`
	RedisHost     string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort     int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewConfig() (Config, error) {
	var conf Config
	if err := env.Parse(&conf); err != nil {
		fmt.Errorf("NewConfig is Error: %+v\n", err)
	}
	return conf, nil
}
