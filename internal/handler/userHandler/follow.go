package userHandler

import (
	"net/http"
	"tiktok/internal/service/userService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
)

type followeMsg struct {
	ToFollowUid string `json:"to_follow_user_id" binding:"required"`
}

func Follow(c *gin.Context) {
	var msg followeMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.Response(c, http.StatusBadRequest, "数据绑定失败", err.Error())
		return
	}

	uid := userService.GetCurUserID(c)

	flag, err := userService.Follow(uid, msg.ToFollowUid)

	if flag == 1 {
		utils.Response(c, http.StatusOK, "关注成功", "")
		return
	} else if flag == 0 {
		utils.Response(c, http.StatusInternalServerError, "关注失败", err.Error())
		return
	} else {
		utils.Response(c, http.StatusOK, "取消关注", "")
		return
	}
}

func CommonFollow(c *gin.Context) {
	var msg followeMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.Response(c, http.StatusBadRequest, "数据绑定失败", err.Error())
		return
	}

	uid := userService.GetCurUserID(c)
	commonFollows, err := userService.CommonFollow(uid, msg.ToFollowUid)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "获取共同关注列表失败", err.Error())
		return
	}

	if len(commonFollows) == 0 {
		utils.Response(c, http.StatusOK, "共同关注为空", "")
		return
	}

	utils.Response(c, http.StatusOK, "获取共同关注成功", commonFollows)
}
