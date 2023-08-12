package core

import (
	"gopkg.in/yaml.v2"
	"gvd_server/config"
	"gvd_server/global"
	"os"
)

const yamlPath = "settings.yaml"

func InitConfig() (c *config.Config) {
	byData, err := os.ReadFile(yamlPath)
	if err != nil {
		global.Log.Fatal("read yaml failed, err:", err)
	}
	c = new(config.Config)
	err = yaml.Unmarshal(byData, c)
	if err != nil {
		global.Log.Fatal("unmarshal yaml failed, err:", err)
	}
	return
}
