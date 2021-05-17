package services

import (
	"github.com/gin-gonic/gin"
)

func ResponseError(e error, statusCode int, c *gin.Context)bool{
	if e!=nil{
		c.JSON(statusCode, gin.H{
			"message": e.Error(),
		})
		c.Abort()
		return true
	}
	return false
}