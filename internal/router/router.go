package router

import (
	"tiktok/internal/handler/userHandler"
	"tiktok/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	u := r.Group("/user")
	u.POST("/register", userHandler.Register)
	u.POST("/login", userHandler.Login)

	return r
}
