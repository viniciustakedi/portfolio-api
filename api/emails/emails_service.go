package emails

import (
	"errors"
	"portfolio/api/config"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailsService struct {
	mongoDB *mongo.Database
}

func NewEmailsService(mongoDB *mongo.Database) *EmailsService {
	return &EmailsService{
		mongoDB: mongoDB,
	}
}

func (ctx *EmailsService) SendPortfolioMessage(data SendPortfolioMessage) (string, error) {
	from := mail.NewEmail("Portfolio Message", "no.reply@takedi.com")
	subject := "New message from portfolio 🎉"
	to := mail.NewEmail("Vinicius Takedi", "viniciustakedi7@gmail.com")

	plainTextContent := getPortfolioMessagePlainText(data)
	htmlContent := getPortfolioMessageHTML(data)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.GetEnv("SENDGRID_API_KEY"))

	response, err := client.Send(message)
	if err != nil {
		return "", err
	}

	if response.StatusCode != 202 {
		return "", errors.New(response.Body)
	}

	return "Email sent successfully!", nil
}
