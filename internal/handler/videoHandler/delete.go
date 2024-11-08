package videoHandler

import (
	"net/http"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/internal/service/userService"
	"tiktok/pkg/utils"
	"tiktok/pkg/utils/minioService"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {

	uid := userService.GetCurUserID(c)

	url := c.PostForm("url")

	db := mysqlDB.GetDB()
	var video mysqlDB.Video
	db.Where("play_url = ?", url).First(&video)
	if video.ID == 0 {
		utils.Response(c, http.StatusBadRequest, "要删除的视频不存在", "")
		return
	} else if video.AuthorID != uid {
		utils.Response(c, http.StatusForbidden, "无删除权限", "")
		return
	}

	if err := minioService.DeleteFileByURL(url); err != nil {
		utils.Response(c, http.StatusInternalServerError, "视频删除失败", err.Error())
		return
	}

	if err := db.Delete(&video).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, "数据库删除视频记录失败", err.Error())
		return
	}

	utils.Response(c, http.StatusOK, "删除成功", "")
}
