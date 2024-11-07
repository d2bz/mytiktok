package videoService

import "tiktok/internal/repository/mysqlDB"

func VideoList() ([]mysqlDB.Video, error) {
	db := mysqlDB.GetDB()
	var cnt int64
	db.Model(&mysqlDB.Video{}).Count(&cnt)

	if cnt > 30 {
		cnt = 30
	}

	vList := make([]mysqlDB.Video, cnt)
	err := db.Model(&mysqlDB.Video{}).Order("ID desc").Limit(int(cnt)).Find(&vList).Error
	return vList, err
}