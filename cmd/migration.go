package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"log"
)

func Migrate(){
	configDb := config.DBEngine.Conn.Config()
	m, err := migrate.New(
		"file://migrations",
		configDb.ConnString())
	if err!=nil{
		log.Printf(configDb.ConnString())
		panic(err)
	}
	m.Steps(2)
}