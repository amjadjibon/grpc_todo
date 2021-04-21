package conf

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type config struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	GrpcPort int `env:"GRPC_PORT" envDefault:"8080"`
	HttpPort int `env:"HTTP_PORT" envDefault:"8090"`
	DSN string `env:"DSN" envDefault:"postgres://amjad:112358@localhost:5432/grpc?sslmode=disable"`
	Debug bool `env:"DEBUG" envDefault:"true"`
}

var Config config

func init() {
	if err := env.Parse(&Config); err != nil {
		log.Err(err)
	}
}

