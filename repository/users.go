package repository

import "context"

type IEUserRepository interface {
	Create(ctx context.Context)
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
