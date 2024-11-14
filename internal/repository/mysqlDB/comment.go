package mysqlDB

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentID string
	UID       string
	VideoID   string
	Content   string
	// 发布时间用gorm.Model里的CreateAt来表示
}

func (*Comment) TableName() string {
	return "comment"
}
