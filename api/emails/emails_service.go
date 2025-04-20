package emails

import (
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

func (ctx *EmailsService) Send() (string, error) {
	return "", nil
}
