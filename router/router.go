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

	// router.Use(func(c *gin.Context) {
	// 	origin := c.GetHeader("Origin")
	// 	if origin == "http://localhost:3000.com" {
	// 		c.Header("Access-Control-Allow-Origin", origin)
	// 		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	// 		if c.Request.Method == "OPTIONS" {
	// 			c.AbortWithStatus(204)
	// 			return
	// 		}
	// 	} else {
	// 		c.AbortWithStatus(403) // Forbidden
	// 		return
	// 	}
	// 	c.Next()
	// })

	api := router.Group("/api")
	{
		RegisterHealthRoutes(api)
<<<<<<< Updated upstream
=======
		RegisterJobsRoutes(api)
		RegisterEmailsRoutes(api)
>>>>>>> Stashed changes
	}

	return router
}
