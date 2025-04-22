package middlewares

import (
	response "portfolio/api/utils"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PayloadValidator(payload interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		payloadInstance := reflect.New(reflect.TypeOf(payload).Elem()).Interface()

		if err := c.ShouldBindJSON(payloadInstance); err != nil {
			response.Error(c, "Error to process JSON: "+err.Error())
			c.Abort()
			return
		}

		validate := validator.New()

		if err := validate.Struct(payloadInstance); err != nil {
			var errors []string
			for _, fieldError := range err.(validator.ValidationErrors) {
				errors = append(errors, fieldError.Error())
			}

			response.Error(c, "Validation failed: "+strings.Join(errors, ", "))
			c.Abort()
			return
		}

		c.Set("payload", payloadInstance)
		c.Next()
	}
}
