package flags

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
)

func DB() {
	fmt.Println("初始化表结构")
	//设定DocModel的tags表为UserModel
	//global.MysqlDB.SetupJoinTable(&models.DocModel{}, "UserCollDocModel", &models.UserModel{})
	err := global.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.RoleModel{},
		&models.DocModel{},
		&models.UserCollDocModel{}, //tags表应在最后生成，否则部分column将被覆盖
		&models.RoleDocModel{},
		&models.ImageModel{},
		&models.UserPwdDocModel{},
		&models.LoginModel{},
		&models.DocDataModel{},
	)
	if err != nil {
		logrus.Fatal("created database autoMigrate failed")
	}
	logrus.Info("数据库迁移成功!")
}
