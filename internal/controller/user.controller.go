package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}
func (uc *UserController) GetUserById(c *gin.Context) {
	user := uc.userService.GetInfoUser("1")

	response.SuccessResponse(c, http.StatusOK, 2001, user);
}
