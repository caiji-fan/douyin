// Package param
// @Author shaofan
// @Date 2022/5/13
// @DESC
package param

// Comment 上传与删除参数
type Comment struct {
	VideoID     int    `json:"video_id"`
	ActionType  byte   `json:"action_type"`
	CommentText string `json:"comment_text"`
	CommentId   string `json:"comment_id"`
}
