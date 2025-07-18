package emails

import (
	"fmt"
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

func (ctx *EmailsController) SendPortfolioMessage(c *gin.Context) {
	data := c.MustGet("payload").(*SendPortfolioMessage)

	message, err := ctx.emailsService.SendPortfolioMessage(*data)
	if err != nil {
		fmt.Println(err.Error())
		response.Error(c, "Error to send email.")
		return
	}

	response.Message(c, message, http.StatusOK)
}

func (ctx *EmailsController) GetNewsletterScheduleTime(newsletterId string) (string, error) {
	time, err := ctx.emailsService.GetNewsletterScheduleTime(newsletterId)
	if err != nil {
		return "", err
	}

	return time, nil
}

func (ctx *EmailsController) SendDailyWordNewsletter() error {
	err := ctx.emailsService.SendDailyWordNewsletter()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (ctx *EmailsController) SendDailyPhrasalVerbNewsletter() error {
	err := ctx.emailsService.SendDailyPhrasalVerbNewsletter()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
