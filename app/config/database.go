package config

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/issengi/goboot/app/services"
)

var DBEngine *DBConnection

type DBConnection struct {
	Conn *pg.DB
	StringConnection string
}

func init() {
	formatSchema := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", Config.DbUser,
		Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName, Config.SSLEnable)
//	DSNConnection := fmt.Sprintf(`user=%s password=%s host=%s port=%s dbname=%s sslmode=%s
//pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s pool_min_conns=%d`, Config.DbUser,
//Config.DbPassword, Config.DbHost, Config.DbPort, Config.DbName, Config.SSLEnable, 500, `30m`, `1m`, 50)
//	configDb, errorParseConfig := pgxpool.ParseConfig(DSNConnection)
//	if errorParseConfig != nil{
//		panic(errorParseConfig)
//	}
//
//	configDb.ConnConfig.Logger = &services.Logger{}
//	db, err := pgxpool.ConnectConfig(context.Background(), configDb)
//	if err!=nil{
//		panic(err)
//	}
//	DBEngine = &DBConnection{Conn: db}
	db := pg.Connect(&pg.Options{
		User: Config.DbUser,
		Password: Config.DbPassword,
		Addr: fmt.Sprintf(`%s:%s`, Config.DbHost, Config.DbPort),
		Database: Config.DbName,
		PoolSize: 50,
		MinIdleConns: 20,
	})
	if Config.IsDev() {
		db.AddQueryHook(services.GoPgLogger{})
	}
	if db == nil{
		panic(errors.New("Failed to connect database"))
	}
	DBEngine = &DBConnection{Conn: db, StringConnection: formatSchema}
}
