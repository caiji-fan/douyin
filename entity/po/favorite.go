// Package po
// @Author shaofan
// @Date 2022/5/13
// @DESC
package po

// Favorite 点赞PO
type Favorite struct {
	RelationModel
	VideoId int `json:"video_id" gorm:"video_id;primaryKey"`
	UserId  int `json:"user_id" gorm:"user_id;primaryKey"`
}
