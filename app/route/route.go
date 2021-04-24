package route

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/actions"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/config"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/middleware"
	"gitlab.com/NeoReids/backend-tryonline-golang/app/services"
)

func init() {
	binding.Validator = new(services.GinDefaultValidator)
}

func InitRoute() {
	baseConfig := config.Config
	if !baseConfig.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowAllOrigins = true
	if baseConfig.CORSAllowOrigins != nil {
		corsConfig.AllowOrigins = baseConfig.CORSAllowOrigins
		corsConfig.AllowAllOrigins = false
	}
	if baseConfig.CORSMethods != nil {
		corsConfig.AllowMethods = baseConfig.CORSMethods
	}
	if baseConfig.CORSHeaders != nil {
		corsConfig.AllowHeaders = baseConfig.CORSHeaders
	}
	router.Use(cors.New(corsConfig))

	router.POST("/login", actions.LoginAction)

	// middleware auth guard
	router.Use(middleware.AuthMiddleware)
	_ = router.Run(fmt.Sprintf(":%s", baseConfig.PortServer))
}