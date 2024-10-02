package routes

import (
	controller "tiktok/controller/UserController"
	"tiktok/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	u := r.Group("/user")
	u.POST("/register", controller.Register)
	u.POST("/login", controller.Login)

	return r
}
