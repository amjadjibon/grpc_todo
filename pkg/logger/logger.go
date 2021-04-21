package logger

import (
	"github.com/amjadjibon/grpc_todo/pkg/conf"
	"github.com/rs/zerolog"
)

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if conf.Config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
