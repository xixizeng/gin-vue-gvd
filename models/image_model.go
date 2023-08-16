package models

import "fmt"

type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	FileName  string    `gorm:"column:fileName;size:128" json:"fileName"`
	Size      int64     `gorm:"column:size;comment:文件大小,单位字节" json:"size"`
	Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`
	Hash      string    `gorm:"column:hash;size:64;comment:文件的hash" json:"hash"`
}

func (image ImageModel) WebPath() string {
	return fmt.Sprintf("%s", image.Path)
}
