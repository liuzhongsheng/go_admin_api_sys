package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goAdminApi/routers/api/v1"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	InitConfing(router)
	SetRouter(router)
	return router
}

// 初始化应用设置
func InitConfing(router *gin.Engine){

}

// 设置路由
func SetRouter(router *gin.Engine){
	api := router.Group("./api")
	{
		v1.RegisterRouter(api)
	}
}