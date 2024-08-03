package service

import "go-ecommerce-backend-api/internal/repo"

type IUserService interface {
	Register(email string, purpose string) int
	GetInfoUser(uid string) string
}

type UserService struct {
	userRepo *repo.UserRepo
}

func (us UserService) Register(email string, purpose string) int {
	//TODO implement me
	panic("implement me")
}

func (us UserService) GetInfoUser(uid string) string {
	//TODO implement me
	panic("implement me")
}

func NewUserService(
	userRepo *repo.UserRepo,
) IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}
