package flashcards

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionCards = "flashcards"
	collectionPaths = "flashcard_paths"
)

type FlashcardsService struct {
	mongoDB *mongo.Database
}

func NewFlashcardsService(mongoDB *mongo.Database) *FlashcardsService {
	return &FlashcardsService{mongoDB: mongoDB}
}

// EnsureIndexes creates compound and text indexes (idempotent).
func (s *FlashcardsService) EnsureIndexes(ctx context.Context) error {
	col := s.mongoDB.Collection(collectionCards)
	_, err := col.Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "language", Value: 1},
				{Key: "path", Value: 1},
				{Key: "difficulty", Value: 1},
			},
		},
		{
			Keys: bson.D{
				{Key: "word", Value: "text"},
				{Key: "description", Value: "text"},
			},
		},
	})
	return err
}

func (s *FlashcardsService) List(ctx context.Context, language, path string, limit, skip int64) ([]FlashcardDocument, error) {
	col := s.mongoDB.Collection(collectionCards)
	filter := bson.M{
		"language": language,
		"path":     path,
	}
	opts := options.Find().
		SetLimit(limit).
		SetSkip(skip).
		SetSort(bson.D{{Key: "difficulty", Value: 1}, {Key: "word", Value: 1}})

	cur, err := col.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []FlashcardDocument
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	if out == nil {
		out = []FlashcardDocument{}
	}
	for i := range out {
		out[i].NormalizeSynonyms()
	}
	return out, nil
}

func (s *FlashcardsService) Count(ctx context.Context, language, path string) (int64, error) {
	col := s.mongoDB.Collection(collectionCards)
	return col.CountDocuments(ctx, bson.M{"language": language, "path": path})
}

func (s *FlashcardsService) GetByID(ctx context.Context, id string) (*FlashcardDocument, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	col := s.mongoDB.Collection(collectionCards)
	var doc FlashcardDocument
	err = col.FindOne(ctx, bson.M{"_id": oid}).Decode(&doc)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	doc.NormalizeSynonyms()
	return &doc, nil
}

func (s *FlashcardsService) ListPaths(ctx context.Context, language string) ([]PathDocument, map[string]int64, error) {
	pcol := s.mongoDB.Collection(collectionPaths)
	cur, err := pcol.Find(ctx, bson.M{"language": language}, options.Find().SetSort(bson.D{{Key: "order", Value: 1}}))
	if err != nil {
		return nil, nil, err
	}
	defer cur.Close(ctx)

	var paths []PathDocument
	if err := cur.All(ctx, &paths); err != nil {
		return nil, nil, err
	}
	if paths == nil {
		paths = []PathDocument{}
	}

	ccol := s.mongoDB.Collection(collectionCards)
	pipe := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"language": language}}},
		{{Key: "$group", Value: bson.M{"_id": "$path", "count": bson.M{"$sum": 1}}}},
	}
	acur, err := ccol.Aggregate(ctx, pipe)
	if err != nil {
		return nil, nil, err
	}
	defer acur.Close(ctx)

	counts := make(map[string]int64)
	for acur.Next(ctx) {
		var row struct {
			ID    string `bson:"_id"`
			Count int64  `bson:"count"`
		}
		if err := acur.Decode(&row); err != nil {
			return nil, nil, err
		}
		counts[row.ID] = row.Count
	}
	if err := acur.Err(); err != nil {
		return nil, nil, err
	}

	return paths, counts, nil
}

func (s *FlashcardsService) Create(ctx context.Context, p *CreateFlashcardPayload) (*FlashcardDocument, error) {
	now := time.Now().UTC()
	doc := FlashcardDocument{
		ID:          primitive.NewObjectID(),
		Word:        p.Word,
		Synonyms:    p.Synonyms,
		Type:        p.Type,
		Language:    p.Language,
		Path:        p.Path,
		Difficulty:  p.Difficulty,
		Description: p.Description,
		Examples:    p.Examples,
		Tags:        p.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	col := s.mongoDB.Collection(collectionCards)
	if _, err := col.InsertOne(ctx, doc); err != nil {
		return nil, err
	}
	return &doc, nil
}

func (s *FlashcardsService) Delete(ctx context.Context, id string) (bool, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid id")
	}
	col := s.mongoDB.Collection(collectionCards)
	res, err := col.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return false, err
	}
	return res.DeletedCount > 0, nil
}
