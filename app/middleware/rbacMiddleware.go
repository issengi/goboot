package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/app/services"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func RbacMiddleware(pathConfig string) gin.HandlerFunc {
	return func(c *gin.Context) {
		mapConfig, errScan := loadFile(pathConfig)
		if services.ResponseError(errScan, http.StatusInternalServerError, c) {
			return
		}
		if statusCode, message := validateRoute(mapConfig, c); statusCode == 200 {
			c.Next()
		} else {
			services.ResponseError(errors.New(message), statusCode, c)
			return
		}
	}
}

func loadFile(path string) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	extension := filepath.Ext(path)
	var errScan error
	// open file json
	byteFile, errOpen := os.Open(path)
	if errOpen!=nil{
		return nil, errOpen
	}
	defer byteFile.Close()
	// read json file parse to interface
	byteValue, _ := ioutil.ReadAll(byteFile)

	switch extension {
	case ".json":
		errScan = json.Unmarshal(byteValue, &jsonMap)
	case ".toml":
		errScan = toml.Unmarshal(byteValue, &jsonMap)
	case ".yml", ".yaml":
		errScan = yaml.Unmarshal(byteValue, &jsonMap)
	}

	return jsonMap, errScan
}

func validateRoute(list map[string]interface{}, c *gin.Context)(statusCode int, message string){
	path := c.FullPath()
	method := c.Request.Method
	routeConfigInterface := services.FindMapInterfaceByKey(list, path)
	if routeConfigInterface != nil {
		methodMap := routeConfigInterface.(map[string]interface{})
		listMiddleware := methodMap[method].([]interface{})
		for _, item := range listMiddleware {
			if !detectHash(item.(string)) {
				isPassed, statusCode := detectRoleFromToken(c, item.(string))
				if !isPassed {
					return statusCode, "Forbidden"
				}
			}
		}
	}
	return http.StatusOK, "ok"
}

func detectRoleFromToken(c *gin.Context, role string) (bool, int) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return false, http.StatusUnauthorized
	}
	token, err := services.JWTParserFromGinContext(c, config.Config.AppKey, "HS256")

	if token != nil && err == nil {
		if _, ok := services.JWTGetClaimsFromToken(token); !ok || !token.Valid {
			return false, http.StatusUnauthorized
		}else{
			mapClaim, _ := services.JWTGetClaimsFromToken(token)
			for _, itemRole := range mapClaim["role"].([]interface{}) {
				if itemRole.(string) == role {
					continue
				}else{
					return false, http.StatusForbidden
				}
			}
			return true, http.StatusOK
		}
	} else {
		return false, http.StatusForbidden
	}
}

func detectHash(text string)(bool){
	if string(text[0]) == `#` && strings.ToLower(string(text[1:])) == "unauthenticated" {
		return true
	}
	return false
}