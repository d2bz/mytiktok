package Redis

import (
	"context"
	"encoding/json"
	"tiktok/common"
	"tiktok/model"
	"time"
)

func SetUserInfo(user model.User) error {
	rdb := common.GetRDB()
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

func GetUserInfo(email string) (model.User, error) {
	rdb := common.GetRDB()
	cbg := context.Background()

	userJson, err := rdb.Get(cbg, email).Result()
	if err != nil {
		return model.User{}, err
	}

	var user model.User
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
