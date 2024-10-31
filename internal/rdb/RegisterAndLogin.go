package rdb

import (
	"context"
	"encoding/json"
	"tiktok/internal/mysqlDB"
	"time"
)

func SetUserInfo(user mysqlDB.User) error {
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

func GetUserInfo(email string) (mysqlDB.User, error) {
	rdb := GetRDB()
	cbg := context.Background()

	userJson, err := rdb.Get(cbg, email).Result()
	if err != nil {
		return mysqlDB.User{}, err
	}

	var user mysqlDB.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return mysqlDB.User{}, err
	}

	return user, nil
}
