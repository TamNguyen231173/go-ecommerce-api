package global

import (
	"github.com/redis/go-redis/v9"
	"go-ecommerce-backend-api/pkg/logger"
	"go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.Zap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
