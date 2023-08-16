package models

// LoginModel 用户登录信息
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	IP        string    `gorm:"size:20" json:"ip"`
	NickName  string    `gorm:"column:nickName;size:42" json:"nickName"`
	UA        string    `gorm:"256" json:"ua"`
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"`
	Addr      string    `gorm:"size:64" json:"addr"`
}
