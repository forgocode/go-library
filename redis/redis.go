package redis

import (
	"github.com/go-redis/redis"
)

func InitRedis(addr, passwd string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		DB:       0,
		Addr:     addr,
		Password: passwd,
		PoolSize: 100,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
