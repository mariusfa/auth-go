package routes

import (
	"github.com/gin-gonic/gin"
	healthController "auth/rest/health/controller"
	userController "auth/rest/user/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Public paths
	r.GET("/health", healthController.GetHealthCheck)
	r.POST("/user", userController.Login)
	
	// Protected paths
	r.Use(userController.AuthMiddleWare())
	r.GET("/protected", userController.ProtectedPath)
	return r
}