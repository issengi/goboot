package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/app/auth"
	"net/http"
)

func LoginAction(c *gin.Context) {
	ruleForm := struct {
		User     string `form:"user" json:"user" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}{}

	if errorValidation := c.ShouldBind(&ruleForm); errorValidation != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": errorValidation.Error(),
		})
	}
	authRepoository := auth.NewAuthRepository()
	token, errorSignToken := authRepoository.CreateJWT(c, ruleForm.User, ruleForm.Password)
	if errorSignToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errorSignToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
