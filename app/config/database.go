package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

var DBEngine *DBConnection

type DBConnection struct {
	Conn *pgx.Conn
}

func init() {
	formatSchema := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", Config.DbUser,
		Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName, Config.SSLEnable)

	db, err := pgx.Connect(context.Background(), formatSchema)
	if err!=nil{
		log.Printf(formatSchema)
		panic(err)
	}
	DBEngine = &DBConnection{Conn: db}
}
