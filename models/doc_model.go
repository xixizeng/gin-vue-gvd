package models

type DocModel struct {
	Model
	Title           string      `gorm:"comment:文档标题;column:title" json:"title"`
	Content         string      `gorm:"comment:文档内容" json:"-"`
	DiggCount       int         `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`
	LookCount       int         `gorm:"comment:浏览量;column:lookCount" json:"lookCount"`
	Key             string      `gorm:"comment:key;not null;unique;column:key" json:"key"`
	ParentID        *uint       `gorm:"comment:父文档id;column:parentID" json:"parentID"`
	ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"` //父文档
	Child           []*DocModel `gorm:"foreignKey:ParentID" json:"-"` //文档的子文档
	FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;JoinForeignKey:DocID;JoinReferences:UserID" json:"-"`
}
