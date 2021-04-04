package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/main/config"
	"github.com/issengi/goboot/main/entities"
	"net/http"
)

func LoginAction(c *gin.Context) {
	ruleForm := struct {
		User		string 		`form:"user" json:"user" binding:"required"`
		Password	string 		`form:"password" json:"password" binding:"required"`
	}{}

	if errorValidation := c.ShouldBind(&ruleForm); errorValidation != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": errorValidation.Error(),
		})
	}

	authRepoository := entities.NewAuthRepository(config.DBEngine)
	user, errorLogin := authRepoository.Login(ruleForm.User, ruleForm.Password)
	if errorLogin!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorLogin.Error(),
		})
		return
	}

	token, errorSignToken := authRepoository.CreateJWT(user)
	if errorSignToken != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorSignToken.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}