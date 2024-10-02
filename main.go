package main

import (
	"tiktok/common"
	"tiktok/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := common.InitDB()
	common.InitRDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic("底层数据库连接失败：" + err.Error())
	}
	defer sqlDB.Close()

	r = routes.Route(r)

	r.Run(":8080")
}
