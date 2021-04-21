package main

import (
	"github.com/amjadjibon/grpc_todo/pkg/pg"
	"github.com/rs/zerolog/log"
)

func main() {
	db := pg.Connect()
	defer db.Close()
	err := pg.CreateSchema(db)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
