package actions

import (
	"github.com/gin-gonic/gin"
	userRepository "github.com/issengi/goboot/users/repository"
	"net/http"
)

func UsersHandler(c *gin.Context) {
	userRepo := userRepository.NewUserRepository()
	total, errorGetCount := userRepo.Count("")
	if errorGetCount != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errorGetCount.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}