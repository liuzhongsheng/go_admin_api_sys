package user

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.RouterGroup)  {
	// 注册
	router.GET("/register", register)
}