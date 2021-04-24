package config

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	//"github.com/go-pg/pg/v10/orm"
)

var DBEngine *pg.DB

func init() {
	DBEngine = pg.Connect(&pg.Options{
		User:     Config.DbUser,
		Password: Config.DbPassword,
		Addr:     fmt.Sprintf("%s:%s", Config.DbHost, Config.DbPort),
		Database: Config.DbName,
	})
}
