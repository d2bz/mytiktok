package router

import (
	"tiktok/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(middleware.CORSMiddleware())

	userRouter(r)
	videoRouter(r)

}
