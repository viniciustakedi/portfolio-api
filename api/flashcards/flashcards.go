package flashcards

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CardType is stored as string in MongoDB.
type CardType string

const (
	TypeVerb         CardType = "verb"
	TypeNoun         CardType = "noun"
	TypeAdjective    CardType = "adjective"
	TypeAdverb       CardType = "adverb"
	TypePhrasalVerb  CardType = "phrasal_verb"
	TypeExpression   CardType = "expression"
)

// FlashcardDocument is the MongoDB shape for flashcards collection.
type FlashcardDocument struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Word         string             `bson:"word" json:"word"`
	Translation  string             `bson:"translation" json:"translation"`
	Type         string             `bson:"type" json:"type"`
	Language     string             `bson:"language" json:"language"`
	Path         string             `bson:"path" json:"path"`
	Difficulty   int                `bson:"difficulty" json:"difficulty"`
	Description  string             `bson:"description" json:"description"`
	Examples     []string           `bson:"examples" json:"examples"`
	Tags         []string           `bson:"tags,omitempty" json:"tags,omitempty"`
	CreatedAt    time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// PathDocument is the MongoDB shape for flashcard_paths collection.
type PathDocument struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Language    string             `bson:"language" json:"language"`
	Level       string             `bson:"level" json:"level"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Order       int                `bson:"order" json:"order"`
	TotalCards  int                `bson:"totalCards" json:"totalCards"`
	Icon        string             `bson:"icon" json:"icon"`
}

// CreateFlashcardPayload is validated JSON body for POST /flashcards.
type CreateFlashcardPayload struct {
	Word        string   `json:"word" validate:"required,min=1,max=200"`
	Translation string   `json:"translation" validate:"required,min=1,max=500"`
	Type        string   `json:"type" validate:"required,oneof=verb noun adjective adverb phrasal_verb expression"`
	Language    string   `json:"language" validate:"required,oneof=en es"`
	Path        string   `json:"path" validate:"required,oneof=beginner intermediate advanced"`
	Difficulty  int      `json:"difficulty" validate:"required,min=1,max=5"`
	Description string   `json:"description" validate:"required,min=1,max=2000"`
	Examples    []string `json:"examples" validate:"required,min=1,dive,required"`
	Tags        []string `json:"tags,omitempty"`
}
