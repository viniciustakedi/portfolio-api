package middlewares

import (
	"net/http"
	"portfolio/api/config"
	response "portfolio/api/utils"

	"github.com/gin-gonic/gin"
)

// RequireFlashcardsAdminKey protects flashcard mutation routes. Set FLASHCARDS_ADMIN_KEY in the environment.
func RequireFlashcardsAdminKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := config.GetEnv("FLASHCARDS_ADMIN_KEY")
		if key == "" {
			response.Error(c, "Flashcards admin is not configured.", http.StatusServiceUnavailable)
			c.Abort()
			return
		}
		if c.GetHeader("X-Admin-Key") != key {
			response.Error(c, "Unauthorized.", http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
