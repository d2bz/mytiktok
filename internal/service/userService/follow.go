package userService

import (
	"context"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"
)

func IsFollow(uid string, toFollowUid string) (bool, error) {
	db := mysqlDB.GetDB()
	var cnt int64
	if err := db.Model(&mysqlDB.Follow{}).
		Where("user_id = ? AND followed_user_id = ?", uid, toFollowUid).
		Count(&cnt).Error; err != nil || cnt == 0 {
		return false, err
	}

	return true, nil
}

func Follow(uid string, toFollowUid string) (int, error) {
	isFollow, err := IsFollow(uid, toFollowUid)
	if err != nil {
		return 0, err
	}

	db := mysqlDB.GetDB()
	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key := redisDB.USER_FOLLOW + uid

	followMsg := &mysqlDB.Follow{
		UserID:         uid,
		FollowedUserID: toFollowUid,
	}

	if !isFollow {
		if err := db.Create(followMsg).Error; err != nil {
			return 0, err
		}

		if _, err := rdb.SAdd(cbg, key, toFollowUid).Result(); err != nil {
			return 0, err
		}

		return 1, nil
	} else {
		//硬删除
		if err := db.Unscoped().
			Where("user_id = ? AND followed_user_id = ?", uid, followMsg.FollowedUserID).
			Delete(followMsg).Error; err != nil {
			return 0, err
		}

		if _, err := rdb.SRem(cbg, key, toFollowUid).Result(); err != nil {
			return 0, err
		}

		return -1, nil
	}
}

func CommonFollow(uid string, targetUid string) ([]string, error) {
	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key1 := redisDB.USER_FOLLOW + uid
	key2 := redisDB.USER_FOLLOW + targetUid

	commonFollows, err := rdb.SInter(cbg, key1, key2).Result()
	if err != nil {
		return nil, err
	}

	return commonFollows, nil
}
