package config

import "fmt"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"`
	DB       string `yaml:"db"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	//下面这俩暂时不要原因未知
	//MaxIdleConnes int    `json:"max-idle-connes" yaml:"max-idle-connes"`
	//MaxOpenConnes int    `json:"max-open-connes" yaml:"max-open-connes"`
	LogLevel string `yaml:"logLevel"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.UserName, m.Password, m.Host, m.Port, m.DB, m.Config)
}

func (m *Mysql) GetLogLevel() string {
	return m.LogLevel
}
