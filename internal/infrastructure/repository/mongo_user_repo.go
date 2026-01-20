package repository

import (
	"context"
	"test-go/internal/features/auth/entity"
	authRepository "test-go/internal/features/auth/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) authRepository.UserRepository {
	return &MongoUserRepository{
		collection: collection,
	}
}

func (m *MongoUserRepository) Create(ctx context.Context, user *entity.User) error {
	_, err := m.collection.InsertOne(ctx, user)
	return err
}

func (m *MongoUserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := m.collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
