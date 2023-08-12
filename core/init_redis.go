package core

import (
	"context"
	"github.com/go-redis/redis"
	"gvd_server/global"
	"time"
)

func InitRedis() *redis.Client {
	redisConf := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password,
		DB:       redisConf.DB,
		PoolSize: redisConf.PoolSize,
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("%s redis failed to connet err:%s", redisConf.Addr(), err)
	}
	return client
}
