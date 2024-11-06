package mysqlDB

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoID  string `gorm:"unique"`
	Title    string
	PlayURL  string `gorm:"unique"`
	AuthorID string
	Liked    int
	IsLiked  bool
}

func (*Video) TableName() string {
	return "video"
}
