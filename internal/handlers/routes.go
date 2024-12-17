package handlers

import (
	"github.com/gin-gonic/gin"
)

// SetupGeneralRoutes SetupRoutes sets up the routes for the application
func SetupGeneralRoutes(router *gin.Engine) *gin.Engine {
	router.GET("/ping", PingHandler)
	return router
}
