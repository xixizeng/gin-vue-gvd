package res

import (
	"github.com/gin-gonic/gin"
	"gvd_server/utils/valid"
	"net/http"
)

type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count int `json:"count"`
}

type Code int

type Response struct {
	Code Code   `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	SUCCESS   = 0
	ErrCode   = 7 //系统错误
	ValidCode = 9 //校验错误
)

func OK(data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: SUCCESS,
		Data: data,
		Msg:  msg,
	})
}

// 响应信息
func OKWithMsg(msg string, c *gin.Context) {
	OK(map[string]any{}, msg, c)
}

// 响应数据
func OKWithData(data any, c *gin.Context) {
	OK(data, "查询成功", c)
}

// 响应详情
func OKWithDetail() {

}

// 响应列表数据
func OKWithList[T any](list []T, count int, c *gin.Context) {
	if len(list) == 0 {
		list = []T{}
	}
	OK(ListResponse[T]{
		List:  list,
		Count: count,
	}, "成功", c)
}

func Fail(code Code, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func FailWithMsg(msg string, c *gin.Context) {
	Fail(ErrCode, map[string]any{}, msg, c)
}

func FailWithValidMsg(msg string, c *gin.Context) {
	Fail(ValidCode, map[string]any{}, msg, c)
}

func FailWithData(data any, c *gin.Context) {
	Fail(ErrCode, data, "系统错误", c)
}

func FailWithError(err error, obj any, c *gin.Context) {
	errorMsg := valid.Error(err)
	Fail(ValidCode, map[string]any{}, errorMsg, c)
}

func FailWithValidError(err error, obj any, c *gin.Context) {
	errorMsg, data := valid.ValidError(err, obj)
	Fail(ValidCode, data, errorMsg, c)
}
