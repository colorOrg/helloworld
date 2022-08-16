package dao

import (
	"context"
	"github.com/go-pg/pg/v10"
)

var Db *pg.DB=InitDb()

func InitDb() *pg.DB {
	Db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "1",
		Database: "demo",

	})
	ctx := context.Background()

	err := Db.Ping(ctx)
	if  err != nil {
		panic(err)
	}
	return Db

}