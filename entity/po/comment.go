// Package po
// @Author shaofan
// @Date 2022/5/13
package po

// Comment 评论PO
type Comment struct {
	EntityModel
	SenderId int    `json:"sender_id" gorm:"sender_id;not null"`
	VideoId  int    `json:"video_id" gorm:"video_id;not null"`
	Content  string `json:"content" gorm:"content;not null"`
	Status   byte   `json:"status" gorm:"status;not null"`
}
