package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/core"
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	global.Log = core.NewLog()
	global.Config = core.InitConfig()
	global.MysqlDB = core.InitMysql()
	global.RedisClient = core.InitRedis()
	//valid.InitTans("zh")

	option := flags.Parse()
	if option.Run() {
		return
	}

	//val, err := global.RedisClient.Get("name1").Result()
	//fmt.Println(val, err)

	logrus.Error("yyy")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	logrus.Info("yyy")
	logrus.Error("yyy")
	global.Log.Info("xxx")
	global.Log.Error("test")

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
