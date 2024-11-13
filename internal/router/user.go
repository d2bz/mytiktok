package router

import (
	"tiktok/internal/handler/userHandler"
	"tiktok/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.Engine) {

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	u := r.Group("/user")
	u.Use(middleware.AuthMiddleware())
	u.POST("/follow", userHandler.Follow)

}
