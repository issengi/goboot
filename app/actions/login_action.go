package actions

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/auth"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
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

	authRepoository := auth.NewAuthRepository(config.DBEngine)
	token, errorSignToken := authRepoository.CreateJWT(ruleForm.User, ruleForm.Password)
	if errorSignToken != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorSignToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
