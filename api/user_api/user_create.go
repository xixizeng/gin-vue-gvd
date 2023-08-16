package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID" binding:"required"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.MysqlDB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		res.FailWithMsg("用户名已存在", c)
		return
	}
	err = global.MysqlDB.Create(&models.UserModel{
		UserName: cr.UserName,
		Password: pwd.HashPwd(cr.Password),
		NickName: cr.NickName,
		RoleID:   cr.RoleID,
		IP:       c.RemoteIP(),
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败", c)
		return
	}

	res.OKWithMsg("成功", c)
	return
}
