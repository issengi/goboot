package main

import (
	"github.com/issengi/goboot/app/config"
	"testing"
)

func TestConnectDb(t *testing.T) {
	db := config.DBEngine
	if db == nil{
		t.Error("Failed to connect database")
	}
}
