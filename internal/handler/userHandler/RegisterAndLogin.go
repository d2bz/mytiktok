package userHandler

import (
	"net/http"
	"regexp"
	"tiktok/internal/mysqlDB"
	"tiktok/internal/rdb"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := mysqlDB.GetDB()

	email := c.PostForm("email")
	password := c.PostForm("password")

	if !isValidForm(c, email, password) {
		return
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
	newUser := mysqlDB.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	db.Create(&newUser)

	//将用户的注册信息存入redis，使注册后五分钟用户的登录不用访问数据库
	err = rdb.SetUserInfo(newUser)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "redis set出错", "")
		return
	}

	utils.Response(c, http.StatusOK, "注册成功", "")
}

func Login(c *gin.Context) {
	db := mysqlDB.GetDB()

	email := c.PostForm("email")
	password := c.PostForm("password")

	if !isValidForm(c, email, password) {
		return
	}

	var user mysqlDB.User
	//尝试从reids中取出用户注册信息，如果不存在再访问数据库
	temUser, err := rdb.GetUserInfo(email)
	if err == nil {
		user = temUser
	} else {
		db.Where("email = ?", email).First(&user)

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

func isValidEmail(email string) bool {
	//使用正则表达式来判断邮箱是否合法
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidForm(c *gin.Context, email string, password string) bool {
	if !isValidEmail(email) {
		utils.Response(c, http.StatusBadRequest, "邮箱格式错误", "")
		return false
	}
	if len(password) < 6 {
		utils.Response(c, http.StatusBadRequest, "密码少于6位", "")
		return false
	}
	return true
}
