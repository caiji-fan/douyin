// Package bo
// @Author shaofan
// @Date 2022/5/14
package bo

const (
	CHANGE_FOLLOW_NUM = '1'
	UPLOAD_VIDEO      = '2'
	FEED_VIDEO        = '3'
)

// RabbitMSG 消息队列标准消息体
type RabbitMSG[T any] struct {
	Type byte `json:"type"`
	// 消息体
	Data T `json:"data"`
	// 重发次数
	ResendCount uint8 `json:"resend_count"`
}
