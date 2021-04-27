package app

import (
	"context"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/route"
)

func InitApp() {
	defer config.DBEngine.Conn.Close(context.Background())
	route.InitRoute()
}
