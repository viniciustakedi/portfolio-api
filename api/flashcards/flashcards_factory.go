package flashcards

import (
	"context"
	"log"
	"sync"
	"time"

	"portfolio/api/infra/db"
)

var indexOnce sync.Once

func MakeFlashcardsController() *FlashcardsController {
	svc := NewFlashcardsService(db.GetMongoDB())
	indexOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := svc.EnsureIndexes(ctx); err != nil {
			log.Printf("flashcards indexes: %v", err)
		}
	})
	return NewFlashcardsController(svc)
}
