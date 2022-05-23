// Package config
// @Author shaofan
// @Date 2022/5/13
package config

// 消息队列相关配置参数
type rabbit struct {
	Url string `yaml:"url"`

	ResendMax int `yaml:"resend-max"`
	Queue     struct {
		ChangeFollowNum string `yaml:"change-follow-num"`
		UploadVideo     string `yaml:"upload-video"`
		FeedVideo       string `yaml:"feed-video"`
	} `yaml:"queue"`

	Key struct {
		ChangeFollowNum string `yaml:"change-follow-num"`
		UploadVideo     string `yaml:"upload-video"`
		FeedVideo       string `yaml:"feed-video"`
	} `yaml:"key"`

	Exchange struct {
		ServiceExchange string `yaml:"service-exchange"`
	} `yaml:"exchange"`
}
