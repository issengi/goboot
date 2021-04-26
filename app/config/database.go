package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DBEngine *gorm.DB

func init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		Config.DbHost, Config.DbUser, Config.DbPassword, Config.DbName, Config.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	DBEngine = db
}
