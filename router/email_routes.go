package router

import (
	"portfolio/api/api/emails"
	"portfolio/api/middlewares"

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
			route:       "/email/send/portfolio-message",
			handler:     emailsController.SendPortfolioMessage,
			middlewares: []gin.HandlerFunc{middlewares.PayloadValidator(&emails.SendPortfolioMessage{})},
		},
	}

	for _, route := range routes {
		switch route.method {
		case "POST":
			router.POST(route.route, append(route.middlewares, route.handler)...)
		}
	}
}
