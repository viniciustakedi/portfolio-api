package emails

type SendPortfolioMessage struct {
	Name    string `json:"name" bson:"name" validate:"required"`
	Email   string `json:"email" bson:"email" validate:"required,email"`
	Message string `json:"message" bson:"message" validate:"required,min=10"`
}
