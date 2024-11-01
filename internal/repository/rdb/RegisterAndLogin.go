package rdb

import (
	"context"
	"encoding/json"
	"tiktok/internal/repository/mysqlDB"

	"time"
)

func SetUserInfoToRedis(user *mysqlDB.User) error {
	rdb := GetRDB()
	cbg := context.Background()

	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = rdb.Set(cbg, user.Email, userJson, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserInfoFromRedis(email string) (*mysqlDB.User, error) {
	rdb := GetRDB()
	cbg := context.Background()

	userJson, err := rdb.Get(cbg, email).Result()
	if err != nil {
		return nil, err
	}

	var user mysqlDB.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
