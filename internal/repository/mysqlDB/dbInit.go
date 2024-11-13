package mysqlDB

import (
	"fmt"
	"tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

// InitDB 开启连接池
func InitDB() *gorm.DB {
	host := config.IP_ADDR
	port := config.DB_PORT
	database := config.DB_DATABASE_NAME
	username := config.DB_USER
	password := config.DB_PASSWORD
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	//自动创建数据表
	createTable(db)
	gormDB = db
	return db
}

func GetDB() *gorm.DB {
	return gormDB
}

func createTable(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&Follow{})
}
