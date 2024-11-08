package videoService

import (
	"context"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func IsVideoLikedByUser(vid string, uid string) (bool, error) {
	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key := redisDB.VIDEO_LIKED + vid

	_, err := rdb.ZScore(cbg, key, uid).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

func LikeVideo(vid string, uid string) error {
	isLike, err := IsVideoLikedByUser(vid, uid)
	if err != nil && err != redis.Nil {
		return err
	}

	db := mysqlDB.GetDB()
	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key := redisDB.VIDEO_LIKED + vid

	if isLike {
		if err := db.Model(&mysqlDB.Video{}).
			Where("video_id = ?", vid).
			Update("liked", gorm.Expr("liked + ?", 1)).
			Error; err != nil {
			return err
		}

		zadd := &redis.Z{
			Score:  float64(time.Now().UnixMilli()),
			Member: uid,
		}

		if _, err := rdb.ZAdd(cbg, key, *zadd).Result(); err != nil {
			return err
		}
	} else {
		if err := db.Model(&mysqlDB.Video{}).
			Where("video_id = ?", vid).
			Update("liked", gorm.Expr("liked - ?", 1)).
			Error; err != nil {
			return err
		}

		if _, err := rdb.ZRem(cbg, key, uid).Result(); err != nil {
			return err
		}
	}

	return nil
}
