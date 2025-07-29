package db

import (
	"context"
	"fmt"
	"portfolio/api/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database

func InitMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := config.GetEnv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("MongoDB connected successfully!")

	config := config.GetEnv("MONGO_DATABASE")
	database = client.Database(config)

	return nil
}

func KillMongoDB() error {
	if database.Client() == nil {
		return fmt.Errorf("mongodb client didn't initialized")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := database.Client().Disconnect(ctx); err != nil {
		return fmt.Errorf("disconnect mongodb: %w", err)
	}
	return nil
}

func GetMongoDB() *mongo.Database {
	return database
}
