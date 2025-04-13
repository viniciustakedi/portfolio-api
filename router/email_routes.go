package router

import (
	"portfolio/api/api/emails"

	"github.com/gin-gonic/gin"
)

func RegisterEmailsRoutes(router *gin.RouterGroup) {
	emailsController := emails.MakeEmailsController()

	routes := []struct {
		method      string
		route       string
		handler     gin.HandlerFunc
		middlewares []gin.HandlerFunc
	}{
		{
			method:      "POST",
			route:       "/email/send",
			handler:     emailsController.Send,
			middlewares: []gin.HandlerFunc{},
		},
	}

	for _, route := range routes {
		switch route.method {
		case "POST":
			router.POST(route.route, append(route.middlewares, route.handler)...)
		}
	}
}
