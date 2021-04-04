package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTParserFromGinContext(c *gin.Context, key, algorithm string) (*jwt.Token, error){
	authorizationHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	return JWTParser(tokenString, key, algorithm)
}

func JWTParser(token, key, algorithm string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(algorithm) != token.Method {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(key), nil
	})
}

func JWTGetClaimsFromToken(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}

