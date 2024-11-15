package videoHandler

import (
	"net/http"
	"tiktok/internal/service/userService"
	"tiktok/internal/service/videoService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
)

type commentMsg struct {
	VideoID string `json:"video_id"`
	Content string `josn:"content"`
}

func PostComment(c *gin.Context) {
	var msg commentMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.Response(c, http.StatusBadRequest, "数据绑定失败", err.Error())
		return
	}

	uid := userService.GetCurUserID(c)

	err := videoService.PostComment(msg.VideoID, uid, msg.Content)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "评论发表失败", err.Error())
		return
	}

	utils.Response(c, http.StatusOK, "评论发表成功", "")
}
