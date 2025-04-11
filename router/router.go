package router

import "github.com/gin-gonic/gin"

func NewRouter(environment string) *gin.Engine {
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil) // Set trusted proxies to nil to disable proxy trust

	api := router.Group("/api")
	{
		RegisterHealthRoutes(api)
	}

	return router
}
