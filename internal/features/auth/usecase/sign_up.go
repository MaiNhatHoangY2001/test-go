package usecase

import (
	"context"
	"test-go/internal/domain/entities"
	"test-go/internal/features/auth/dto"
	"test-go/internal/features/auth/repository"
	errs "test-go/internal/shared/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignUpUseCase struct {
	userRepo repository.UserRepository
}

func NewSignUpUseCase(userRepo repository.UserRepository) *SignUpUseCase {
	return &SignUpUseCase{
		userRepo: userRepo,
	}
}

func (uc *SignUpUseCase) Execute(ctx context.Context, req *dto.SignUpRequest) (*dto.AuthResponse, error) {
	existingUser, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to check user existence")
	}
	if existingUser != nil {
		return nil, errs.New(errs.ConflictError, "User already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.Wrap(err, errs.InternalError, "Failed to hash password")
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
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to create user")
	}

	token, err := GenerateJWT(user.Email, user.ID.Hex())
	if err != nil {
		return nil, errs.Wrap(err, errs.InternalError, "Failed to generate token")
	}

	return &dto.AuthResponse{
		Token: token,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
