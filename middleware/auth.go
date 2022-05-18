// Package middleware
// @Author shaofan
// @Date 2022/5/13
package middleware

import (
	"douyin/util/jwtutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// JWTAuth jwt鉴权
func JWTAuth(ctx *gin.Context) {

	//获取参数user_id
	userId, err1 := strconv.Atoi(ctx.Query("user_id"))
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "未传入user_id",
		})
		ctx.Abort()
		return
	}
	tokenString := ctx.Query("token")
	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "未传入token",
		})
		//停止
		ctx.Abort()
		return
	}
	//解析出token中的用户id
	uid, err2 := jwtutil.ParseJWT(tokenString)
	if err2 != nil {
		//token过期
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "token已过期",
		})
		ctx.Abort()
		return
	}
	//TODO 查询redis中是否存在token并和此token对比是否一致
	//_, err3 := redis.GetString(tokenString)
	//if err3 != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{
	//		"status_code": http.StatusBadRequest,
	//		"status_msg":  "redis中token不存在",
	//	})
	//	ctx.Abort()
	//	return
	//}

	//对比两个id是否一致
	if uid != userId {
		//id不一致
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "无权限",
		})
		ctx.Abort()
		return
	}
	//继续执行下面的程序
	ctx.Next()
}
