package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CloseMongoClient(ctx context.Context, client *mongo.Client) error {
	return client.Disconnect(ctx)
}

// EnsureIndexes creates database indexes for optimal query performance
func EnsureIndexes(ctx context.Context, db *mongo.Database) error {
	// Create indexes for todos collection
	todosCollection := db.Collection("todos")
	
	// Index on created_at for sorting (descending for newest first)
	_, err := todosCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: map[string]interface{}{"created_at": -1},
		Options: options.Index().SetName("idx_created_at"),
	})
	if err != nil {
		return err
	}

	// Create indexes for users collection
	usersCollection := db.Collection("users")
	
	// Unique index on email for fast lookups and uniqueness
	_, err = usersCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]interface{}{"email": 1},
		Options: options.Index().SetUnique(true).SetName("idx_email_unique"),
	})
	if err != nil {
		return err
	}

	return nil
}
