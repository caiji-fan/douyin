// Package param
// @Author shaofan
// @Date 2022/5/13
package param

// Favorite 点赞参数
type Favorite struct {
	UserID     int  `json:"user_id" form:"user_id"`
	VideoID    int  `json:"video_id" form:"video_id"`
	ActionType byte `json:"action_type" form:"action_type"  binding:"required" msg:"无效的操作类型"`
}
