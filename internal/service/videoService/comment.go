package videoService

import (
	"context"
	"encoding/json"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// 发表评论
// 创建评论信息 -> 存入数据库 -> 存入缓存
func PostComment(vid string, uid string, content string) error {
	cid, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	comment := &mysqlDB.Comment{
		CommentID: cid.String(),
		UID:       uid,
		VideoID:   vid,
		Content:   content,
	}

	commentJson, err := json.Marshal(comment)
	if err != nil {
		return err
	}

	db := mysqlDB.GetDB()
	err = db.Create(comment).Error
	if err != nil {
		return err
	}

	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key1 := redisDB.COMMENT_INFO + comment.CommentID

	err = rdb.Set(cbg, key1, commentJson, -1).Err()
	if err != nil {
		return err
	}

	zadd := &redis.Z{
		Score:  float64(time.Now().UnixMilli()),
		Member: comment.CommentID,
	}

	key2 := redisDB.VIDEO_COMMENT + vid
	err = rdb.ZAdd(cbg, key2, *zadd).Err()
	return err
}
