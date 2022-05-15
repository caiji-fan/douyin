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
func InitRoute() {
	route := gin.Default()

	withAUTH := route.Group("douyin", middleware.JWTAuth)
	user1 := withAUTH.Group("user")
	{
		user1.GET("", controller.UserInfo)
	}
	publish := withAUTH.Group("publish")
	{
		publish.GET("list", controller.VideoList)
		publish.POST("action", controller.Publish)
	}
	favorite := withAUTH.Group("favorite")
	{
		favorite.GET("list", controller.FavoriteList)
		favorite.POST("action", controller.Like)
	}
	comment := withAUTH.Group("comment")
	{
		comment.GET("list", controller.CommentList)
		comment.POST("action", controller.Comment)
	}
	relation := withAUTH.Group("relation")
	{
		relation.GET("follow/list", controller.FollowList)
		relation.GET("follower/list", controller.FansList)
		relation.POST("action", controller.Follow)
	}

	withoutAUTH := route.Group("douyin")
	withoutAUTH.GET("feed", controller.Feed)
	user2 := withoutAUTH.Group("user")
	{
		user2.POST("register", controller.Register)
		user2.POST("login", controller.Login)
	}
}
