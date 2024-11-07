package userHandler

import (
	"log"
	"net/http"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/repository/rdb"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	db := mysqlDB.GetDB()

	email := c.PostForm("email")
	password := c.PostForm("password")

	//数据验证
	if matched, err := utils.IsValidForm(utils.EmailPattern, email); !matched {
		if err == nil {
			utils.Response(c, http.StatusBadRequest, "邮箱格式错误", "")
			return
		} else {
			utils.Response(c, http.StatusInternalServerError, "手机号格式匹配出错", err.Error())
			return
		}
	}

	if matched, err := utils.IsValidForm(utils.PasswordPattern, password); !matched {
		if err == nil {
			utils.Response(c, http.StatusBadRequest, "密码格式错误", "")
			return
		} else {
			utils.Response(c, http.StatusInternalServerError, "密码格式匹配出错", err.Error())
			return
		}
	}

	//var user mysqlDB.User
	//尝试从reids中取出用户注册信息，如果不存在再访问数据库
	user, err := rdb.GetUserInfoFromRedis(email)
	if err != nil || user == nil {
		log.Printf("err: %v", err)
		err = db.Where("email = ?", email).First(&user).Error //使用&user而不是user，user可能为nil，使用&user gorm自动为user进行内存分配
		if err != nil {
			utils.Response(c, http.StatusInternalServerError, "登录时从数据库获取用户信息失败", err.Error())
			return
		}

		if user.ID == 0 {
			utils.Response(c, http.StatusBadRequest, "用户不存在", "")
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		utils.Response(c, http.StatusBadRequest, "密码错误", "")
		return
	}

	token, err := utils.ReleaseToken(user)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "token生成失败", "")
		return
	}

	utils.Response(c, http.StatusOK, "登录成功", token)
}
