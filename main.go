package main

import (
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/app/route"
)

func main() {
	defer config.DBEngine.Close()
	route.InitRoute()
}