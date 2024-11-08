package main

import (
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/redisDB"
	"tiktok/internal/router"
	"tiktok/pkg/utils/minioService"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := mysqlDB.InitDB()
	redisDB.InitRDB()
	minioService.MinioInit()

	sqlDB, err := db.DB()
	if err != nil {
		panic("底层数据库连接失败：" + err.Error())
	}
	defer sqlDB.Close()

	router.Router(r)

	r.Run(":8080")
}
