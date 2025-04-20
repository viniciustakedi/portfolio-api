package emails

import (
	"net/http"
	response "portfolio/api/utils"

	"github.com/gin-gonic/gin"
)

type EmailsController struct {
	emailsService *EmailsService
}

func NewEmailsController(
	emailsService *EmailsService,
) *EmailsController {
	return &EmailsController{
		emailsService: emailsService,
	}
}

func (ctx *EmailsController) Send(c *gin.Context) {
	message, err := ctx.emailsService.Send()
	if err != nil {
		response.Error(c, "Error to send email.")
		return
	}

	response.Message(c, message, http.StatusOK)
}
