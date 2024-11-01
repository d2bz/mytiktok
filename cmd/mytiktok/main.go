package main

import (
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/rdb"
	"tiktok/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := mysqlDB.InitDB()
	rdb.InitRDB()

	sqlDB, err := db.DB()
	if err != nil {
		panic("底层数据库连接失败：" + err.Error())
	}
	defer sqlDB.Close()

	r = router.Route(r)

	r.Run(":8080")
}
