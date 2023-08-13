package models

type UserCollDocModel struct {
	Model
	DocID     uint      `gorm:"column:docID" json:"docID"`
	DocModel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
