package routers

import (
	"go-ecommerce-backend-api/internal/routers/manager"
	"go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserGroupRouter
	Manager manager.AdminGroupRouter
}

var RouterGroupApp = new(RouterGroup)
