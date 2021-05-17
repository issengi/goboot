package route

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/issengi/goboot/app/middleware"
	userRepository "github.com/issengi/goboot/users/actions"
	//"github.com/gin-gonic/gin/binding"
	"github.com/issengi/goboot/app/actions"
	"github.com/issengi/goboot/app/config"
	//"github.com/issengi/goboot/app/services"
)

//func init() {
//	binding.Validator = new(services.GinDefaultValidator)
//}

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
	router.Use(middleware.RbacMiddleware(baseConfig.RouteConfig))
	router.POST("/login", actions.LoginAction)

	// middleware auth guard
	//router.Use(middleware.AuthMiddleware)
	router.GET("/total-user/:id", userRepository.UsersHandler)
	_ = router.Run(fmt.Sprintf(":%s", baseConfig.PortServer))
}
