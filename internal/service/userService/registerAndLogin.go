package userService

import (
	"context"
	"encoding/json"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"

	"time"
)

func SetLoginUser(user *mysqlDB.User) error {
	rdb := redisDB.GetRDB()
	cbg := context.Background()
	key := redisDB.LOGIN_USER_EMAIL + user.Email

	userJson, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = rdb.Set(cbg, key, userJson, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetLoginUser(email string) (*mysqlDB.User, error) {
	rdb := redisDB.GetRDB()
	cbg := context.Background()
	key := redisDB.LOGIN_USER_EMAIL + email

	userJson, err := rdb.Get(cbg, key).Result()
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
