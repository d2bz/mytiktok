package videoHandler

import (
	"net/http"
	"tiktok/internal/service/userService"
	"tiktok/internal/service/videoService"
	"tiktok/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetVideoList(c *gin.Context) {

	uid := userService.GetCurUserID(c)

	vList, err := videoService.VideoList(uid)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "获取视频列表失败", err.Error())
		return
	}

	utils.Response(c, http.StatusOK, "获取视频列表成功", vList)
}
