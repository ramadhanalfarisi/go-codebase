package drivers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/ramadhanalfarisi/go-codebase/config"
)


var ctx = context.Background()

func RedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     config.REDIS_URL,
		Password: config.REDIS_PASSWORD,
		DB:       0,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(pong)
	return client
}

func SetRedisValue(key string, value string) bool {
	client := RedisConnection()
	err := client.Set(ctx, key, value, 1*time.Minute).Err()
	if err != nil {
		log.Println(err)
	}
	client.Close()
	return true
}

func GetRedisValue(key string) string {
	client := RedisConnection()
	get, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	client.Close()
	return get
}

func DeleteRedisValue(keys []string) bool {
	client := RedisConnection()
	err := client.Del(ctx, keys...).Err()
	if err != nil {
		log.Println(err)
	}
	client.Close()
	return true
}

func SearchRedisValue(keys string) []string {
	client := RedisConnection()
	search, err := client.Keys(ctx, keys).Result()
	if err != nil {
		log.Println(err)
	}
	client.Close()
	return search
}
