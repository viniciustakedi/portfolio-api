package flashcards

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ErrAlreadySeeded means flashcard_paths already has documents and force was false.
var ErrAlreadySeeded = errors.New("flashcard data already exists; run with force to replace")

// SeedCollections inserts default paths and flashcards. With force=true, clears both collections first.
func SeedCollections(ctx context.Context, database *mongo.Database, force bool) error {
	pcol := database.Collection(collectionPaths)
	ccol := database.Collection(collectionCards)

	if force {
		if _, err := pcol.DeleteMany(ctx, bson.M{}); err != nil {
			return fmt.Errorf("clear paths: %w", err)
		}
		if _, err := ccol.DeleteMany(ctx, bson.M{}); err != nil {
			return fmt.Errorf("clear flashcards: %w", err)
		}
	} else {
		n, err := pcol.CountDocuments(ctx, bson.M{})
		if err != nil {
			return err
		}
		if n > 0 {
			return ErrAlreadySeeded
		}
	}

	svc := NewFlashcardsService(database)
	if err := svc.EnsureIndexes(ctx); err != nil {
		return fmt.Errorf("indexes: %w", err)
	}

	paths := seedPaths()
	pdocs := make([]interface{}, len(paths))
	for i := range paths {
		pdocs[i] = paths[i]
	}
	if _, err := pcol.InsertMany(ctx, pdocs); err != nil {
		return fmt.Errorf("insert paths: %w", err)
	}

	cards := seedCards()
	now := time.Now().UTC()
	cdocs := make([]interface{}, len(cards))
	for i := range cards {
		cards[i].CreatedAt = now
		cards[i].UpdatedAt = now
		cdocs[i] = cards[i]
	}
	if _, err := ccol.InsertMany(ctx, cdocs); err != nil {
		return fmt.Errorf("insert flashcards: %w", err)
	}

	// Sync totalCards on path documents from live counts
	for _, lang := range []string{"en", "es"} {
		for _, lvl := range []string{"beginner", "intermediate", "advanced"} {
			cnt, err := ccol.CountDocuments(ctx, bson.M{"language": lang, "path": lvl})
			if err != nil {
				return err
			}
			_, err = pcol.UpdateOne(ctx,
				bson.M{"language": lang, "level": lvl},
				bson.M{"$set": bson.M{"totalCards": cnt}},
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
