// Package route
// @Author shaofan
// @Date 2022/5/13
package route

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

// InitRoute 初始化接口路由
func InitRoute() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.Default()

	//路由分组信息
	//relation
	relationGroup := router.Group("/douyin/relation")
	{
		//关注/取消关注
		relationGroup.POST("/action", middleware.JWTAuth, controller.Follow)
		//用户关注列表
		relationGroup.GET("/follow/list", middleware.JWTAuth, controller.FollowList)
		//用户粉丝列表
		relationGroup.GET("/follower/list", middleware.JWTAuth, controller.FansList)

	}
	//favorite
	favoriteGroup := router.Group("/douyin/favorite")
	{
		//关注/取消关注
		favoriteGroup.POST("/action", middleware.JWTAuth, controller.Like)
		//用户关注列表
		favoriteGroup.GET("/list", middleware.JWTAuth, controller.FavoriteList)

	}
	return router

}
