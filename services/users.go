package services

type IEUserService interface {
	Create()
}

type UserService struct{}

func (u *UserService) Create() {

}

func NewUserService() *UserService {
	return &UserService{}
}
