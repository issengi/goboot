package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strings"
)

const (
	DEV = "dev"
	LOCAL = "local"
)

var (
	Config *config
)

type config struct {
	AppName				string
	AppKey				string
	// database config
	DbHost				string
	DbUser 				string
	DbPassword 			string
	DbName				string
	SSLEnable			string
	DbPort				string

	// port server address
	PortServer			string

	// mode main LOCAL, DEV, PRODUCTION
	AppEnv				string

	// CORS config
	CORSAllowOrigins	[]string
	CORSMethods			[]string
	CORSHeaders			[]string
}

func init() {
	// load .env file with error return
	e := godotenv.Load()
	if e != nil{
		panic(".env file not found.")
	}
	Config = &config{}
	// list map string for env to struct key format {ENV_KEY: STRUCT_KEY}
	listEnv := map[string]string{
		"DB_HOST": "DbHost",
		"DB_NAME": "DbName",
		"DB_USER": "DbUser",
		"DB_PASSWORD": "DbPassword",
		"DB_PORT": "DbPort",
		"SSL_DB": "SSLEnable",
		"PORT_SERVER": "PortServer",
	}
	for key, v := range listEnv {
		Config.Manipulate(v, os.Getenv(key))
	}
	devMode := os.Getenv("APP_ENV")
	if devMode == "" {
		devMode = DEV
	}
	Config.Manipulate("AppEnv", devMode)

	appKey := "thisisasecretkeyforjwt!@#"
	if os.Getenv("APP_KEY") != "" {
		appKey = os.Getenv("APP_KEY")
	}
	Config.Manipulate("AppKey", appKey)

	appName := "local_app"
	if os.Getenv("APP_NAME") != "" {
		appName = os.Getenv("APP_NAME")
	}
	Config.Manipulate("AppName", appName)

	listCORSAllowOrigin := strings.Split(os.Getenv("CORS_ALLOW_ORIGINS"),",")
	if os.Getenv("CORS_ALLOW_ORIGINS") != ""{
		fmt.Println(len(listCORSAllowOrigin))
		Config.CORSAllowOrigins = listCORSAllowOrigin
	}
	if os.Getenv("CORS_METHODS") != "" {
		Config.CORSMethods = strings.Split(os.Getenv("CORS_METHODS"), ",")
	}
	if os.Getenv("CORS_HEADERS") != ""	{
		Config.CORSHeaders = strings.Split(os.Getenv("CORS_HEADERS"), ",")
	}
}

func (c *config) Manipulate(key string, value string){
	reflection := reflect.ValueOf(c)
	fieldname := reflect.Indirect(reflection).FieldByName(key)
	if fieldname.Kind() != reflect.Invalid {
		fieldname.SetString(value)
	}
}

func (c *config) IsDev() bool {
	return c.AppEnv == DEV || c.AppEnv == LOCAL
}