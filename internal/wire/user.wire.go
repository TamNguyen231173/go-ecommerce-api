package wire

import (
	"github.com/google/wire"
	"go-ecommerce-backend-api/internal/controller"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		controller.NewUserController,
		service.NewUserService,
		repo.NewUserRepository())
	return new(controller.UserController), nil
}
