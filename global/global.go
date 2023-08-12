package global

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvd_server/config"
)

var (
	Config      *config.Config   //配置文件
	Log         *logrus.Logger   //日志设定
	MysqlDB     *gorm.DB         //数据库DB
	MysqlLog    logger.Interface //日志等级
	RedisClient *redis.Client
)
