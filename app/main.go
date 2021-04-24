package app

import (
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/route"
)

func InitApp() {
	defer config.DBEngine.Close()
	route.InitRoute()
}
