package initialize

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/routers"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r := gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
	}
	//middleware
	//r.Use() //logging
	//r.Use() //cross
	//r.Use() //limiter global
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/checkStatus") // tracking monitoring
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitAdminRouter(MainGroup)
	}
	return r
}
