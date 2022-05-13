// Package config
// @Author shaofan
// @Date 2022/5/13
package config

// 消息队列相关配置参数
type rabbit struct {
	Url string `yaml:"url"`

	Queue struct {
		ChangeFollowNum string `yaml:"change-follow-num"`
		UploadVideo     string `yaml:"upload-video"`
		FeedVideo       string `yaml:"feed-video"`
	} `yaml:"queue"`

	exchange struct {
		ChangeFollowNum string `yaml:"change-follow-num"`
		UploadVideo     string `yaml:"upload-video"`
		FeedVideo       string `yaml:"feed-video"`
	} `yaml:"exchange"`

	Key struct {
		ServiceExchange string `yaml:"dy_exchange"`
	} `yaml:"key"`
}
