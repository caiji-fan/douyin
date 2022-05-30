// Package bo
// @Author shaofan
// @Date 2022/5/31
package bo

type ChangeFollowNumBody struct {
	UserId   int  `json:"user_id"`
	ToUserId int  `json:"to_user_id"`
	IsFollow bool `json:"is_follow"`
}
