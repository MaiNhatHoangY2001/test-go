package repository

import (
	"context"
	"errors"
	"test-go/internal/features/todo/entity"
	todoRepository "test-go/internal/features/todo/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTodoRepository struct {
	collection *mongo.Collection
}

func NewMongoTodoRepository(collection *mongo.Collection) todoRepository.TodoRepository {
	return &MongoTodoRepository{
		collection: collection,
	}
}

func (m *MongoTodoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	_, err := m.collection.InsertOne(ctx, todo)
	return err
}

func (m *MongoTodoRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	result, err := m.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("todo not found")
	}

	return nil
}

func (m *MongoTodoRepository) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	cursor, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var todos []*entity.Todo
	if err = cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (m *MongoTodoRepository) GetByID(ctx context.Context, id string) (*entity.Todo, error) {
	var todo entity.Todo

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	err = m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}

	return &todo, nil
}

func (m *MongoTodoRepository) Update(ctx context.Context, todo *entity.Todo) error {
	_, err := m.collection.ReplaceOne(
		ctx,
		bson.M{"_id": todo.ID},
		todo,
	)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("todo not found")
		}
		return err
	}
	return nil
}
