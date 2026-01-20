package usecase

import (
	"context"
	"test-go/internal/features/auth/dto"
	"test-go/internal/features/auth/repository"
	errs "test-go/internal/shared/errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to find user")
	}
	if user == nil {
		return nil, errs.New(errs.BadRequestError, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errs.New(errs.BadRequestError, "Invalid credentials")
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
