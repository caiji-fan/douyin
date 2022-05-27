// Package bo
// @Author shaofan
// @Date 2022/5/22
package bo

import "time"

// Feed 视频流存储实体
type Feed struct {
	VideoId    int       `json:"video_id"`
	CreateTime time.Time `json:"create_date"`
}
