package mysqlDB

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserID         string
	FollowedUserID string
}

func (*Follow) TableName() string {
	return "follow"
}
