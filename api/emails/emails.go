package emails

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SendPortfolioMessage struct {
	Name    string `json:"name" bson:"name" validate:"required"`
	Email   string `json:"email" bson:"email" validate:"required,email"`
	Message string `json:"message" bson:"message" validate:"required,min=10"`
}

type SendDailyWordNewsletter struct {
	Name    string `json:"name" bson:"name" validate:"required"`
	Email   string `json:"email" bson:"email" validate:"required,email"`
	Message string `json:"message" bson:"message" validate:"required,min=10"`
}

type WordInfo struct {
	Word       string   `json:"word"`
	Definition string   `json:"definition"`
	UsageTip   string   `json:"usageTip"`
	FunFact    string   `json:"funFact"`
	Examples   []string `json:"examples"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type RecipientsDB struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email            string             `bson:"email" json:"email"`
	Name             string             `bson:"name" json:"name"`
	SubscribedAt     time.Time          `bson:"subscribedAt" json:"subscribedAt"`
	NewsletterTypeId primitive.ObjectID `bson:"newsletterTypeId" json:"newsletterTypeId"`
}

type WordDB struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Used bool               `bson:"used"`
}
