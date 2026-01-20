package repository

import (
"context"
"test-go/internal/features/auth/entity"
)

type UserRepository interface {
Create(ctx context.Context, user *entity.User) error
FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
