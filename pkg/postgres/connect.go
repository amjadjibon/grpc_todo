package postgres

import (
	"github.com/amjadjibon/grpc_todo/pkg/conf"
	"github.com/go-pg/pg/v10"
	"log"
)

func Connect() *pg.DB {
	opt, err := pg.ParseURL(conf.Config.DSN)
	if err != nil {
		log.Println(err)
	}
	db := pg.Connect(opt)
	return db
}

