package pwd

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPwd 将密码哈希
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logrus.Error(err)
	}
	return string(hash)
}

// CheckPwd 验证密码 hashPwd hash后的密码 需要被验证的密码
func CheckPwd(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
