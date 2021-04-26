package cmd

import (
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/domain"
)

func Migrate() error{
	db := config.DBEngine
	return db.AutoMigrate(&domain.Users{}, &domain.Roles{})
}