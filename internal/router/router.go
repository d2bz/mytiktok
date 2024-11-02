package router

import (
	"tiktok/internal/handler/userHandler"
	videohandler "tiktok/internal/handler/videoHandler"
	"tiktok/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	u := r.Group("/user")
	u.POST("/register", userHandler.Register)
	u.POST("/login", userHandler.Login)

	v := r.Group("/video")
	v.POST("/publish", videohandler.Publish)

	return r
}
