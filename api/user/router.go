package user

import (
	"firstGin/middleware/cors"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	r := e.Group("/api/v1")
	//使用中间件
	r.Use(cors.TestMiddle())
	r.POST("/userLogin", userLogin)
}
