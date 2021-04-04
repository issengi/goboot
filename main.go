package main

import (
	"github.com/issengi/goboot/main/config"
	"github.com/issengi/goboot/main/route"
)

func main() {
	defer config.DBEngine.Close()
	route.InitRoute()
}