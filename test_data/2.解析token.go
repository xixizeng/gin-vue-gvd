package main

import (
	"fmt"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/utils/jwts"
)

func main() {
	//claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrTmFtZSI6InN4eCIsInJvbGVJRCI6MiwidXNlcklEIjoxLCJleHAiOjE2OTE5Nzg3MjkuNDc3NjkxLCJpc3MiOiJ4aXhpIn0.BGCYZyf19Nj9xxLIFsiseI6RidUJpZ9iKTOjpMMs_Ck")
	//if err != nil {
	//	logrus.Fatal("token 无效")
	//}
	//fmt.Println(claims)
	global.Config = core.InitConfig()
	token, err := jwts.GenToken(jwts.JwyPayLoad{
		NickName: "xioxi",
	})
	cliams, err := jwts.ParseToken(token)
	fmt.Println(cliams, err)
}
