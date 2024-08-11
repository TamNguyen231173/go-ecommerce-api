package user

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/internal/wire"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()

	//public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/login", userController.Login)
	}
	//private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.POST("/get-info", userController.GetUserById)
	}
}
