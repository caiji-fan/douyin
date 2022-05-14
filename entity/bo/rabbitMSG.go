// Package bo
// @Author shaofan
// @Date 2022/5/14
package bo

// RabbitMSG 消息队列标准消息体
type RabbitMSG[T any] struct {
	// 消息体
	Data T `json:"data"`
	// 重发次数
	ResendCount uint8 `json:"resend_count"`
}
