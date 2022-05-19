// Package middleware
// @Author shaofan
// @Date 2022/5/13
package middleware

import (
	"douyin/config"
	"douyin/util/jwtutil"
	"douyin/util/redisutil"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// JWTAuth jwt鉴权
func JWTAuth(ctx *gin.Context) {
	var err error
	var tokenString string
	var userId int
	//判断投稿接口
	path := ctx.Request.URL.Path
	flag := strings.Contains(path, "/douyin/publish/action")
	if flag {
		//投稿接口 通过form-data获取token
		tokenString = ctx.PostForm("token")
	} else {
		//获取参数userId
		userId, err = getUserId(ctx)
		if err != nil {
			return
		}
		//获取token
		tokenString = ctx.Query("token")
	}

	//解析token
	err, uid := parseToken(ctx, tokenString)
	if err != nil {
		return
	}
	//从redis判断token是否有效
	err = tokenValid(ctx, uid)
	if err != nil {
		return
	}
	//对比两个id是否一致 (投稿接口不需要此判断)
	if !flag {
		err = equalId(ctx, userId, uid)
		if err != nil {
			return
		}
	}
	//投稿接口存入uid，供线程变量中间件操作
	ctx.Set(config.Config.ThreadLocal.Keys.UserId, uid)
	//继续执行下面的程序
	ctx.Next()
}

//获取参数user_Id
func getUserId(ctx *gin.Context) (int, error) {
	userId, err := strconv.Atoi(ctx.Query("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "未传入user_id",
		})
		ctx.Abort()
	}
	return userId, err
}

//解析token
func parseToken(ctx *gin.Context, token string) (error, int) {
	//解析出token中的用户id
	uid, err := jwtutil.ParseJWT(token)
	if err != nil {
		//token过期
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "token已过期/未获取到token参数",
		})
		ctx.Abort()
	}
	return err, uid
}

//判断redis中token是否有效
func tokenValid(ctx *gin.Context, userId int) error {
	var redisToken string
	err := redisutil.Get(config.Config.Redis.Key.Token+strconv.Itoa(userId), &redisToken)
	if err != nil || redisToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "redis中token不存在",
		})
		ctx.Abort()
		return errors.New("redis中token不存在")
	}
	return err
}

//判断两个id是否一致
func equalId(ctx *gin.Context, userId int, uId int) error {
	if userId != uId {
		//id不一致
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "无权限",
		})
		ctx.Abort()
		return errors.New("id不一致")
	}
	return nil
}
