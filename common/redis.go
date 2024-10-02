package common

import "github.com/redis/go-redis/v9"

var Rdb *redis.Client

func InitRDB() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "124.71.229.101:6379",
		Password: "tiktok",
		DB:       0,
	})

	Rdb = rdb
}

func GetRDB() *redis.Client {
	return Rdb
}
