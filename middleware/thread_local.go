// Package middleware
// @Author shaofan
// @Date 2022/5/18
package middleware

import (
	"douyin/config"
	"github.com/gin-gonic/gin"
	"github.com/timandy/routine"
)

var ThreadLocal = routine.NewInheritableThreadLocal()

// SaveUserId 存放用户Id
func SaveUserId(ctx *gin.Context) {
	var storage = make(map[string]string, 1)
	var userId = ctx.Keys[config.Config.ThreadLocal.Keys.UserId]
	if userId != nil {
		storage[config.Config.ThreadLocal.Keys.UserId] = userId.(string)
		ThreadLocal.Set(storage)
		ctx.Next()
		ThreadLocal.Remove()
	}
}