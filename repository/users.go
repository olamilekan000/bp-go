package repository

type IEUserRepository interface {
	Create()
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
