// Package param
// @Author shaofan
// @Date 2022/5/13
package param

// User 用户注册与登录参数
type User struct {
	UserName string `form:"username" binding:"required" msg:"无效的用户名"`
	Password string `form:"password"  binding:"gte=6" msg:"密码格式不正确"`
}
