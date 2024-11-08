package userService

import (
	// "errors"
	"tiktok/internal/repository/mysqlDB"
	// "tiktok/internal/repository/redisDB"

	"github.com/gin-gonic/gin"
)

// func GetUserInfoByID(uid string) (*mysqlDB.User, error) {
// 	user, err := redisDB.GetUserInfoFromRedis(uid)
// 	if err != nil || user == nil {
// 		db := mysqlDB.GetDB()
// 		db.Where("UID = ?", uid).First(&user)

// 		if user.ID == 0 {
// 			return nil, errors.New("用户不存在")
// 		}
// 	}

// 	return user, nil
// }

func GetCurUserID(c *gin.Context) string {
	c_uid, _ := c.Get("curUser")
	uid := c_uid.(mysqlDB.User).UID
	return uid
}
