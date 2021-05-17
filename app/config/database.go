package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DBEngine *DBConnection

type DBConnection struct {
	Conn 				*sql.DB
	StringConnection 	string
	Log					*log.Logger
}

func init() {
	formatSchema := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", Config.DbUser,
		Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName, Config.SSLEnable)

	stringConnection := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`, Config.DbHost,
		Config.DbPort, Config.DbUser, Config.DbPassword, Config.DbName, Config.SSLEnable)

	db, err := sql.Open("postgres", stringConnection)
	if err!=nil{
		log.Println(err.Error())
		os.Exit(1)
	}

	DBEngine = &DBConnection{
		Conn: db,
		StringConnection: formatSchema,
		Log: log.New(os.Stdout, `[SQL]`, log.LstdFlags),
	}
}

func (d DBConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
	d.Log.Println(query, args)
	return d.Conn.Exec(query, args...)
}

func (d DBConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	d.Log.Println(query, args)
	return d.Conn.Query(query, args...)
}

func (d DBConnection) QueryRow(query string, args ...interface{}) *sql.Row {
	d.Log.Println(query, args)
	return d.Conn.QueryRow(query, args...)
}