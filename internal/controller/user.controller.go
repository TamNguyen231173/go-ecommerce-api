package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"net/http"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	user := uc.userService.Register("email", "purpose")
	response.SuccessResponse(c, http.StatusOK, 2001, user)
}
