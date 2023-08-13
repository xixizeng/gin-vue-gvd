package models

type RoleModel struct {
	Model
	Title    string     `gorm:"size:16;not null;comment:角色名称" json:"title"`                                    //角色名称
	Pwd      string     `gorm:"size:32" json:"-"`                                                              //角色密码
	IsSystem bool       `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               //是否是系统角色
	DocsList []DocModel `gorm:"many2many:role_doc_models;JoinForeignKey:RoleID;JoinReferences:DocID" json:"-"` //角色拥有的文档列表
}
