package models

type UserModel struct {
	Model
	UserName  string    `gorm:"colum:userName;comment:用户名;size:36;unique;not null" json:"userName"` //用户名
	Password  string    `gorm:"colum:password;comment:密码;size:128;" json:"password"`                //密码
	Avatar    string    `gorm:"colum:avatar;comment:头像;size:256" json:"avatar"`                     //头像
	NickName  string    `gorm:"colum:nickName;comment:昵称;size:36" json:"nickName"`                  //昵称
	Email     string    `gorm:"colum:email;comment:邮箱;size:128" json:"email"`                       //邮箱
	Token     string    `gorm:"colum:token;comment:其他平台的唯一ID;size:256" json:"token"`                //其他平台的唯一ID
	IP        string    `gorm:"colum:ip;comment:ip;size:16" json:"ip"`                              //ip
	Addr      string    `gorm:"colum:addr;comment:地址;size:64" json:"addr""`                         //地址
	RoleID    uint      `gorm:"colum:roleID;comment:用户对应的角色" json:"roleID"`                         //用户对应的角色
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"roleModel"`
}
