// Package route
// @Author shaofan
// @Date 2022/5/13
package route

import (
	"douyin/controller"
	"github.com/gin-gonic/gin"
)

// InitRoute 初始化接口路由
func InitRoute(r *gin.Engine) {
	router := r.Group("/douyin/user")
	{
		router.POST("/register/", controller.UserControllers{}.Register)
		router.POST("/login/", controller.UserControllers{}.Login)
		router.GET("/", controller.UserControllers{}.UserInfo)
	}
}
