package videoService

import (
	"context"
	"encoding/json"
	"log"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"
)

func CommentList(vid string) ([]mysqlDB.Comment, error) {
	//尝试从缓存中获取
	comments, err := getCommentListFromRDB(vid)
	if comments != nil {
		return comments, nil
	}

	//缓存没有就从数据库里查
	log.Println("Error getting comments from RDB:", err)
	db := mysqlDB.GetDB()
	err = db.Model(&mysqlDB.Comment{}).
		Where("video_id = ?", vid).
		Order("ID desc").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func getCommentListFromRDB(vid string) ([]mysqlDB.Comment, error) {
	comments := make([]mysqlDB.Comment, 0)

	cbg := context.Background()
	rdb := redisDB.GetRDB()
	key := redisDB.VIDEO_COMMENT + vid
	commentZStructs, err := rdb.ZRevRangeWithScores(cbg, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	//存入redis的是comment的Json字符串，要把字符串转换为结构体实例
	for _, v := range commentZStructs {
		var comment mysqlDB.Comment
		err := json.Unmarshal(v.Member.([]byte), &comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
