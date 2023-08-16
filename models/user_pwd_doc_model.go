package models

type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"userID"json:"userID"`
	DocID  uint `gorm:"docID" json:"docID"`
}
