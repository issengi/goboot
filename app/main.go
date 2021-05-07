package app

import (
	"context"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/app/route"
)

func InitApp() {
	defer config.DBEngine.Conn.Close(context.Background())
	route.InitRoute()
}
