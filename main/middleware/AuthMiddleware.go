package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/main/config"
	"github.com/issengi/goboot/main/services"
	"net/http"
	"strings"
)

func AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Token",
		})
		c.Abort()
		return
	}
	token, err := services.JWTParserFromGinContext(c, config.Config.AppKey, "HS256")

	if token != nil && err == nil {
		if _, ok := services.JWTGetClaimsFromToken(token); !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is already expired.",
			})
			c.Abort()
		}
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
}