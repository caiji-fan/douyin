// Package douyin
// @Author shaofan
// @Date 2022/5/13
// @DESC
package main

import (
	"douyin/config"
	"douyin/repositories"
)

func main() {
}

func init() {
	// 配置文件初始化
	config.Init()
	// 数据库初始化
	repositories.Init()
}
