// Package bo
// @Author shaofan
// @Date 2022/5/31
package bo

type RabbitErrorMSG struct {
	FeedVideo       []RabbitMSG[int]                 `json:"feed_video"`
	UploadVideo     []RabbitMSG[int]                 `json:"upload_video"`
	ChangeFollowNum []RabbitMSG[ChangeFollowNumBody] `json:"change_follow_num"`
}
