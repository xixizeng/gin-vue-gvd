package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey"`             //主键ID
	CreatedAt time.Time      `gorm:"column:createdAt"json:"createdAt"` //添加时间
	UpdateAt  time.Time      `gorm:"column:updateAt"json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
