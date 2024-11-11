package userService

import (
	//"context"
	"tiktok/internal/repository/mysqlDB"
	//"tiktok/internal/repository/redisDB"
)

func IsFollow(uid string, toFollowUid string) {

}

func Follow(uid string, toFollowUid string) {
	IsFollow(uid, toFollowUid)

	db := mysqlDB.GetDB()
	// cbg := context.Background()
	// rdb := redisDB.GetRDB()
	// key

	followMsg := &mysqlDB.Follow{
		UserID:         uid,
		FollowedUserID: toFollowUid,
	}

	if true {
		if err := db.Create(followMsg).Error; err != nil {
			return
		}

		// msg存入redis
	} else {
		//改成硬删除
		if err := db.Delete(followMsg).Error; err != nil {
			return
		}

		// 从set删除取关的博主
	}
}
