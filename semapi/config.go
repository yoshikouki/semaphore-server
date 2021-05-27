package semapi

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/imdario/mergo"
)

type Config struct {
	Port          int    `env:"PORT" envDefault:"8686"`
	RedisHost     string `env:"REDIS_HOST" envDefault:"localhost"`
	RedisPort     int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewConfig(conf Config) (Config, error) {
	var defaultConfig Config
	if err := env.Parse(&defaultConfig); err != nil {
		fmt.Errorf("NewConfig is Error: %+v\n", err)
	}

	mergo.Merge(&conf, defaultConfig)
	return conf, nil
}
