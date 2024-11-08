package redisDB

import (
	"tiktok/config"

	"github.com/redis/go-redis/v9"
)

var redisDB *redis.Client

func InitRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RDB_ADDR,
		Password: config.RDB_PASSWORD,
		DB:       0,
	})

	redisDB = rdb
}

func GetRDB() *redis.Client {
	return redisDB
}
