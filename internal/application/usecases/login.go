package usecases

import (
	"context"
	"errors"
	"test-go/internal/application/dto"
	"test-go/internal/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	userRepo repositories.UserRepository
}

func NewLoginUseCase(userRepo repositories.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		userRepo: userRepo,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
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
