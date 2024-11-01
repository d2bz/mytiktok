package service

import (
	"errors"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/rdb"
)

func GetUserInfoByID(uid string) (*mysqlDB.User, error) {
	user, err := rdb.GetUserInfoFromRedis(uid)
	if err != nil || user == nil {
		db := mysqlDB.GetDB()
		db.Where("UID = ?", uid).First(user)

		if user.ID == 0 {
			return nil, errors.New("用户不存在")
		}
	}

	return user, nil
}
