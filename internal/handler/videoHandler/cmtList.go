package videoHandler

import (
	"net/http"
	"tiktok/internal/service/videoService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
)

type videoMsg struct {
	VideoID string `json:"video_id" bind:"required"`
}

func GetCommentList(c *gin.Context) {
	var msg videoMsg
	if err := c.ShouldBindJSON(&msg); err != nil {
		utils.Response(c, http.StatusBadRequest, "数据绑定失败", err.Error())
		return
	}

	comments, err := videoService.CommentList(msg.VideoID)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "获取评论列表失败", err.Error())
		return
	}

	utils.Response(c, http.StatusOK, "获取评论列表成功", comments)
}
