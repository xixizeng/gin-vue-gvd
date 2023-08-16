package main

import (
	"fmt"
	"gvd_server/utils/pwd"
)

func main() {
	hash := pwd.HashPwd("1234")
	hash1 := pwd.HashPwd("1234")
	fmt.Println(hash, hash1)
}
