// Package param
// @Author shaofan
// @Date 2022/5/13
package param

const (
	DO_COMMENT     = 1
	DELETE_COMMENT = 0
)

// Comment 上传与删除参数
type Comment struct {
	VideoID     int    `json:"video_id"`
	ActionType  byte   `json:"action_type"`
	CommentText string `json:"comment_text"`
	UserId      int    `json:"user_id"`
	CommentId   int    `json:"comment_id"`
}
