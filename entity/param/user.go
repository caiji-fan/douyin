// Package param
// @Author shaofan
// @Date 2022/5/13
package param

// User 用户注册与登录参数
type User struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
