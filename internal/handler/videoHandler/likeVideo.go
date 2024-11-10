package videoHandler

import (
	"net/http"
	"tiktok/internal/service/userService"
	"tiktok/internal/service/videoService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
)

type likeMsg struct {
	//Uid string `json:"user_id" binding:"required"`
	Vid string `json:"video_id" binding:"required"`
}

func LikeVideo(c *gin.Context) {
	var msg likeMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.Response(c, http.StatusBadRequest, "数据绑定失败", err.Error())
		return
	}

	uid := userService.GetCurUserID(c)

	flag, err := videoService.LikeVideo(msg.Vid, uid)

	if flag == 1 {
		utils.Response(c, http.StatusOK, "点赞成功", "")
		return
	} else if flag == 0 {
		utils.Response(c, http.StatusInternalServerError, "点赞失败", err.Error())
		return
	} else {
		utils.Response(c, http.StatusOK, "取消点赞", "")
		return
	}
}
