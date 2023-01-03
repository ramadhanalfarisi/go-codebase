package helpers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/ramadhanalfarisi/go-codebase-kocak/config"
)

var redis_host,
	redis_port,
	redis_pass,
	redis_addr, environment string

var ctx = context.Background()

func init() {
	environment = "development"
}

func RedisConnection() *redis.Client {
	if environment == "test" {
		redis_host = config.REDIS_HOST_TEST
		redis_port = config.REDIS_PORT_TEST
		redis_pass = config.REDIS_PASSWORD_TEST
	} else if environment == "development" {
		redis_host = config.REDIS_HOST_DEV
		redis_port = config.REDIS_PORT_DEV
		redis_pass = config.REDIS_PASSWORD_DEV
	} else {
		redis_host = config.REDIS_HOST_PROD
		redis_port = config.REDIS_PORT_PROD
		redis_pass = config.REDIS_PASSWORD_PROD
	}

	redis_addr = fmt.Sprintf("%s:%s", redis_host, redis_port)
	client := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: redis_pass,
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
