package router

import (
	"tiktok/internal/handler/videoHandler"
	"tiktok/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func videoRouter(r *gin.Engine) {
	v := r.Group("/video")
	v.Use(middleware.AuthMiddleware())
	v.POST("/publish", videoHandler.Publish)
	v.POST("/delete", videoHandler.Delete)
	v.GET("/list", videoHandler.GetVideoList)
	v.POST("/like", videoHandler.LikeVideo)

	cmt := v.Group("/comment")
	cmt.POST("/post", videoHandler.PostComment)
	cmt.GET("/list", videoHandler.GetCommentList)

}
