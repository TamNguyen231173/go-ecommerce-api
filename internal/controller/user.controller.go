package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/service"
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

	c.JSON(http.StatusOK, gin.H{
		"message": "this is user controller",
		"user":    user,
	})
}
