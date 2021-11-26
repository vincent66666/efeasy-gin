package initialize

import (
	"context"
	"efeasy-gin/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"strconv"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Default.Host + ":" + strconv.Itoa(global.App.Config.Redis.Default.Port),
		Password: global.App.Config.Redis.Default.Password, // no password set
		DB:       global.App.Config.Redis.Default.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Logger.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
