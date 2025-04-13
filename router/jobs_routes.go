package router

import (
	"portfolio/api/api/jobs"

	"github.com/gin-gonic/gin"
)

func RegisterJobsRoutes(router *gin.RouterGroup) {
	jobsController := jobs.MakeJobsController()

	routes := []struct {
		method      string
		route       string
		handler     gin.HandlerFunc
		middlewares []gin.HandlerFunc
	}{
		{
			method:      "GET",
			route:       "/jobs",
			handler:     jobsController.GetAll,
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
