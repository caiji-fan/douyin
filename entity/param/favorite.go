// Package param
// @Author shaofan
// @Date 2022/5/13
package param

// Favorite 点赞参数
type Favorite struct {
	UserID     int  `json:"user_id"`
	VideoID    int  `json:"video_id"`
	ActionType byte `json:"action_type"`
}
