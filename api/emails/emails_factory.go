package emails

import "portfolio/api/infra/db"

func MakeEmailsController() *EmailsController {
	emailsService := NewEmailsService(db.GetMongoDB())
	emailsController := NewEmailsController(emailsService)

	return emailsController
}
