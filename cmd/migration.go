package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/issengi/goboot/app/config"
	"log"
)

func Migrate(){
	configDb := config.DBEngine
	m, err := migrate.New(
		"file://migrations",
		configDb.StringConnection)
	if err!=nil{
		log.Printf(configDb.StringConnection)
		panic(err)
	}
	if errStep := m.Steps(12); errStep != nil{
		log.Println(errStep.Error())
	}
}