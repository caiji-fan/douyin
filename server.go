// Package douyin
// @Author shaofan
// @Date 2022/5/13
// @DESC
package main

import (
	"douyin/config"
	"douyin/repositories/daoimpl"
	"douyin/route"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	r.Run()
}

func init() {
	// 配置文件初始化
	config.Init()
	// 数据库初始化
	daoimpl.Init()

	//路由初始化
	r = route.InitRoute()
}
