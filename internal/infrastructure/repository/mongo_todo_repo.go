package repository

import (
	"context"
	"test-go/internal/domain/entities"
	errs "test-go/internal/shared/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTodoRepository struct {
	collection *mongo.Collection
}

func NewMongoTodoRepository(collection *mongo.Collection) *MongoTodoRepository {
	return &MongoTodoRepository{
		collection: collection,
	}
}

func (m *MongoTodoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	_, err := m.collection.InsertOne(ctx, todo)
	return err
}

func (m *MongoTodoRepository) Delete(ctx context.Context, userID, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errs.New(errs.BadRequestError, "Invalid todo ID format")
	}

	result, err := m.collection.DeleteOne(ctx, bson.M{"_id": objID, "user_id": userID})
	if err != nil {
		return errs.Wrap(err, errs.DatabaseError, "Failed to delete todo")
	}

	if result.DeletedCount == 0 {
		return errs.New(errs.NotFoundError, "Todo not found")
	}

	return nil
}

func (m *MongoTodoRepository) GetAll(ctx context.Context, userID string, page, limit int) ([]*entities.Todo, int64, error) {
	// Filter by user ID
	filter := bson.M{"user_id": userID}
	
	// Get total count for this user
	totalCount, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Calculate skip value
	skip := (page - 1) * limit

	// Find with pagination
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	
	cursor, err := m.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var todos []*entities.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, 0, err
	}

	return todos, totalCount, nil
}

func (m *MongoTodoRepository) GetByID(ctx context.Context, userID, id string) (*entities.Todo, error) {
	var todo entities.Todo

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.New(errs.BadRequestError, "Invalid todo ID format")
	}

	err = m.collection.FindOne(ctx, bson.M{"_id": objID, "user_id": userID}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.New(errs.NotFoundError, "Todo not found")
		}
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to retrieve todo")
	}

	return &todo, nil
}

func (m *MongoTodoRepository) Update(ctx context.Context, userID string, todo *entities.Todo) error {
	result, err := m.collection.ReplaceOne(
		ctx,
		bson.M{"_id": todo.ID, "user_id": userID},
		todo,
	)
	if err != nil {
		return errs.Wrap(err, errs.DatabaseError, "Failed to update todo")
	}
	
	if result.MatchedCount == 0 {
		return errs.New(errs.NotFoundError, "Todo not found")
	}
	
	return nil
}
