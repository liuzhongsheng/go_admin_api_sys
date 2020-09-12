package v1

import (
	"github.com/gin-gonic/gin"
	"goAdminApi/routers/api/v1/user"
)

func RegisterRouter(router *gin.RouterGroup){
	v1 := router.Group("/v1")
	{
		// 用户路由
		user.RegisterRouter(v1.Group("/user"))
	}
}