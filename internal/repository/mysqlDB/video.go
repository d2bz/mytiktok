package mysqlDB

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	VideoID  string `gorm:"unique"`
	Title    string
	PlayURL  string `gorm:"unique"`
	AuthorID string
	Liked    int
}

func (*Video) TableName() string {
	return "video"
}

type apiVideo struct {
	Title    string
	PlayURL  string
	AutherID string
	liked    int
	IsLiked  bool
}
