package mysqlDB

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	VideoID  string
	Title    string
	PlayURL  string
	AuthorID string
	Liked    int
	IsLiked  bool
}

func (*Video) TableName() string {
	return "video"
}
