package api

import (
	"example/Controllers"
	v1 "example/pkg/services/v1/onboarding"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// default gin router engine which have inbuilt functionality for logger and recovery
	router := gin.Default()
	flag := false
	var login v1.LogInterface
	login = v1.LogController{}

	v1 := router.Group("v1")
	{
		onboard := v1.Group("onboard")
		{
			onboard.POST("/registration", login.Registration)
			onboard.POST("/login", login.Login)
			onboard.GET("/logout", login.Logout)
		}
		user := v1.Group("/user")
		{
			user.GET("/", Controllers.ListUser)
			user.POST("/", Controllers.AddNewUser)
			user.GET("/:id", Controllers.GetOneUser)
			user.PUT("/:id", Controllers.PutOneUser)
			user.DELETE("/:id", Controllers.DeleteUser)
		}

	}
	// set true want to use default router
	//flag = true
	if flag {
		return router
	}

	router1 := gin.New()
	router1.Use(gin.Logger())
	router1.Use(gin.Recovery())
	v2 := router1.Group("v2")
	{
		onboard := v2.Group("onboard")
		{
			onboard.POST("/registration", login.Registration)
			onboard.POST("/login", login.Login)
			onboard.GET("/logout", login.Logout)
		}
		user := v2.Group("/user")
		{
			user.GET("/", Controllers.ListUser)
			user.POST("/", Controllers.AddNewUser)
			user.GET("/:id", Controllers.GetOneUser)
			user.PUT("/:id", Controllers.PutOneUser)
			user.DELETE("/:id", Controllers.DeleteUser)
		}
	}

	return router1
}
