package router

import (
	"portfolio/api/api/health"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(router *gin.RouterGroup) {
	healthController := health.MakeHealthController()

	routes := []struct {
		method      string
		route       string
		handler     gin.HandlerFunc
		middlewares []gin.HandlerFunc
	}{
		{
			method:      "GET",
			route:       "/health",
			handler:     healthController.HealthCheck,
			middlewares: []gin.HandlerFunc{},
		},
	}

	for _, route := range routes {
		switch route.method {
		case "GET":
			router.GET(route.route, append(route.middlewares, route.handler)...)
		}
	}
}
