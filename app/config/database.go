package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DBEngine *DBConnection

type DBConnection struct {
	Conn 				*sqlx.DB
	StringConnection 	string
	Log					*log.Logger
}

func init() {
	formatSchema := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", Config.DbUser,
		Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName, Config.SSLEnable)

	stringConnection := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`, Config.DbHost,
		Config.DbPort, Config.DbUser, Config.DbPassword, Config.DbName, Config.SSLEnable)


	db, err := sqlx.Connect("postgres", stringConnection)
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

func (d DBConnection) InsertReturnId(query string, arg interface{}) (int64, error) {
	var result int64
	rows, err := d.Conn.NamedQuery(query, arg)
	defer rows.Close()
	if err!=nil{
		return result, err
	}

	for rows.Next() {
		rows.Scan(&result)
	}

	return result, nil
}

//func (d DBConnection) Exec(query string, args ...interface{}) (sql.Result, error) {
//	d.Log.Println(query, args)
//	return d.Conn.Exec(query, args...)
//}
//
//func (d DBConnection) Query(query string, args ...interface{}) (*sql.Rows, error) {
//	d.Log.Println(query, args)
//	return d.Conn.Query(query, args...)
//}
//
//func (d DBConnection) QueryRow(query string, args ...interface{}) *sql.Row {
//	d.Log.Println(query, args)
//	return d.Conn.QueryRow(query, args...)
//}