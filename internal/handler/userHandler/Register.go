package userHandler

import (
	"net/http"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/service/userService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
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

	var user mysqlDB.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		utils.Response(c, http.StatusBadRequest, "邮箱已被注册", "")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "密码加密出错", "")
		return
	}

	name := "用户" + utils.RandomString(10)
	uid, err := uuid.NewUUID()
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "id生成错误", "")
		return
	}
	newUser := mysqlDB.User{
		UID:      uid.String(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	db.Create(&newUser)

	//将用户的注册信息存入redis，使注册后五分钟用户的登录不用访问数据库
	err = userService.SetLoginUser(&newUser)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "redis set出错", "")
		return
	}

	utils.Response(c, http.StatusOK, "注册成功", "")
}
