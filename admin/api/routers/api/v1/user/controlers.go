package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary 用户注册
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200
// @Failure 500
// @Router /api/v1/tags [post]
func register(c *gin.Context) {
fmt.Print("111")
}