package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: "",
		DB:       0,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
	}

	fmt.Println("Connected to Redis")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		fmt.Printf("Failed to set key: %s\n", err)
		return
	}
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Printf("Failed to get key: %s\n", err)
		return
	}
	global.Logger.Info("value score is::", zap.String("value", value))
}
