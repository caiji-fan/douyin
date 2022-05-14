// Package po
// @Author shaofan
// @Date 2022/5/14
// @DESC
package po

// EntityModel 实体模型
type EntityModel struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	CreateTime string `json:"create_time" gorm:"create_time;not null"`
	UpdateTime string `json:"update_time" gorm:"update_time;not null"`
}

// RelationModel 关系模型
type RelationModel struct {
	CreateTime string `json:"create_time" gorm:"create_time;not null"`
	UpdateTime string `json:"update_time" gorm:"update_time;not null"`
}
