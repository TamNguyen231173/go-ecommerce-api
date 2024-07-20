package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/login")
	}
	//private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.POST("/get-info")
	}
}
