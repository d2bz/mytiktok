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

type ApiVideo struct {
	VideoID  string
	Title    string
	PlayURL  string
	AutherID string
	liked    int
	IsLiked  bool
}

func (v *Video) ToApiVideo() (*ApiVideo) {

	av := &ApiVideo{
		VideoID:  v.VideoID,
		Title:    v.Title,
		PlayURL:  v.PlayURL,
		AutherID: v.AuthorID,
		liked:    v.Liked,
		IsLiked:  false,
	}

	return av
}
