package videoService

import (
	"tiktok/internal/repository/mysqlDB"
)

func VideoList(uid string) (*[]mysqlDB.ApiVideo, error) {
	db := mysqlDB.GetDB()
	var cnt int64
	db.Model(&mysqlDB.Video{}).Count(&cnt)

	if cnt > 30 {
		cnt = 30
	}

	vList := make([]mysqlDB.Video, cnt)
	err := db.Model(&mysqlDB.Video{}).Order("ID desc").Limit(int(cnt)).Find(&vList).Error
	if err != nil {
		return nil, err
	}

	avList := make([]mysqlDB.ApiVideo, 0, cnt)
	for _, v := range vList {
		av := v.ToApiVideo()

		isLiked, err := IsVideoLikedByUser(v.VideoID, uid)
		if err != nil {
			return nil, err
		}

		av.IsLiked = isLiked

		avList = append(avList, *av)
	}
	return &avList, nil
}
