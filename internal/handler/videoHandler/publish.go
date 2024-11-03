package videoHandler

import (
	"net/http"
	"tiktok/internal/repository/mysqlDB"
	"tiktok/pkg/utils"
	"tiktok/pkg/utils/minioService"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Publish(c *gin.Context) {
	url, err := minioService.UploadVideo(c)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "文件上传出错", err.Error())
		return
	}

	title := c.PostForm("videoTitle")

	c_uid, _ := c.Get("curUser")
	uid := c_uid.(mysqlDB.User).UID

	vid, err := uuid.NewUUID()
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "vid生成错误", err.Error())
		return
	}

	video := mysqlDB.Video{
		VideoID:  vid.String(),
		Title:    title,
		PlayURL:  url,
		AuthorID: uid,
		Liked:    0,
		IsLiked:  false,
	}

	db := mysqlDB.GetDB()
	db.Create(&video)
	utils.Response(c, http.StatusOK, "视频上传成功", url)
}
