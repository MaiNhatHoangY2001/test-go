package usecases

import (
	"context"
	"errors"
	"test-go/internal/application/dto"
	"test-go/internal/domain/entities"
	"test-go/internal/domain/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignUpUseCase struct {
	userRepo repositories.UserRepository
}

func NewSignUpUseCase(userRepo repositories.UserRepository) *SignUpUseCase {
	return &SignUpUseCase{
		userRepo: userRepo,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, req *dto.SignUpRequest) (*dto.AuthResponse, error) {
	existingUser, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:        primitive.NewObjectID(),
		Email:     req.Email,
		Password:  string(hashedPassword),
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := GenerateJWT(user.Email, user.ID.Hex())
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
