package global

import (
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Mdb    *gorm.DB
)
