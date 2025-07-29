package router

import (
	"slices"

	"github.com/gin-gonic/gin"
)

func NewRouter(environment string) *gin.Engine {
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	router.Use(func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		originsAllowed := []string{
			"https://takedi.com",
			"https://www.takedi.com",
			"https://my-portfolio-git-v2-viniciustakedis-projects.vercel.app",
		}

		if environment == "development" {
			originsAllowed = append(originsAllowed, origin)
		}

		if slices.Contains(originsAllowed, origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
		} else {
			c.AbortWithStatus(403) // Forbidden
			return
		}
		c.Next()
	})

	api := router.Group("/api")
	{
		RegisterHealthRoutes(api)
		RegisterJobsRoutes(api)
		RegisterEmailsRoutes(api)
	}

	return router
}
