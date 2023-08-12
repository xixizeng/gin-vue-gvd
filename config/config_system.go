package config

import "fmt"

type System struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (this System) Addr() string {
	return fmt.Sprintf("%s:%d", this.IP, this.Port)
}
