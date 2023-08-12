package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvd_server/global"
	"time"
)

func Gorm() *gorm.DB {
	return InitMysql()
}

func InitMysql() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		logrus.Warn("未配置mysql,取消连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var logLevel logger.LogLevel
	switch global.Config.Mysql.LogLevel {
	case "info":
		logLevel = logger.Info
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	default:
		logLevel = logger.Info //默认打印INfo
	}
	global.MysqlLog = logger.Default.LogMode(logLevel)

	//这一段是根据环境设定数据库错误输出等级
	//if global.Config.System.Env == "dev" {
	//	global.MysqlLog = logger.Default.LogMode(logger.Info)
	//} else {
	//	global.MysqlLog = logger.Default.LogMode(logger.Error)
	//}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   global.MysqlLog,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		global.Log.Error(fmt.Sprintf("[%s] mysql connet failed err:%s", dsn, err))
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大连接数
	sqlDB.SetMaxIdleConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间,不能草果mysql的wait_timeout
	return db
}
