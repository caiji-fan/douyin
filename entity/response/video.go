// Package response
// @Author shaofan
// @Date 2022/5/19
package response

import "douyin/entity/bo"

// 视频这一块的返回应该是啥？单个的视频？视频列表？

type VideoList struct {
	Response
	Data []bo.Video `json:"video_list"`
}

type Video struct {
	Response
	Data bo.Video `json:"video"`
}
