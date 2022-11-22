package services

import "context"

type IEUserService interface {
	Create(ctx context.Context)
}

type UserService struct{}

func (u *UserService) Create(ctx context.Context) {

}

func NewUserService() *UserService {
	return &UserService{}
}
