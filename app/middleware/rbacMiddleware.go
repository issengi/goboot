package middleware

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/app/config"
	"github.com/issengi/goboot/app/services"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func RbacMiddleware(pathConfig string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// open file json
		jsonFile, err := os.Open(pathConfig)
		defer jsonFile.Close()
		if services.ResponseError(err, http.StatusInternalServerError, c) {
			return
		}

		// read json file parse to interface
		byteValue, _ := ioutil.ReadAll(jsonFile)
		jsonMap := make(map[string]interface{})
		err = json.Unmarshal(byteValue, &jsonMap)
		if services.ResponseError(err, http.StatusInternalServerError, c) {
			return
		}
		if statusCode, message := ValidateRoute(jsonMap, c); statusCode == 200 {
			c.Next()
		} else {
			services.ResponseError(errors.New(message), statusCode, c)
			return
		}
	}
}

func ValidateRoute(list map[string]interface{}, c *gin.Context)(statusCode int, message string){
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