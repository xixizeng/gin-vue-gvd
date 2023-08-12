package flags

import (
	"fmt"
	"gvd_server/global"
)

func Port(port int) {
	global.Config.System.Port = port
	fmt.Println("修改程序运行端口")
}
