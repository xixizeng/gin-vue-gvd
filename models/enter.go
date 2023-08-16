package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	gorm.Model
	ID        uint           `json:"id" gorm:"primaryKey"`             //主键ID
	CreatedAt time.Time      `gorm:"column:createdAt"json:"createdAt"` //添加时间
	UpdatedAt time.Time      `gorm:"column:updatedAt"json:"updateAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
