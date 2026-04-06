package router

import (
	"portfolio/api/api/flashcards"
	"portfolio/api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterFlashcardsRoutes(router *gin.RouterGroup) {
	ctrl := flashcards.MakeFlashcardsController()
	admin := middlewares.RequireFlashcardsAdminKey()

	router.GET("/flashcards", ctrl.List)
	router.GET("/flashcards/paths", ctrl.ListPaths)
	router.POST("/flashcards", middlewares.PayloadValidator(&flashcards.CreateFlashcardPayload{}), admin, ctrl.Create)
	router.GET("/flashcards/:id", ctrl.GetByID)
	router.DELETE("/flashcards/:id", admin, ctrl.Delete)
}
