package valid

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Trans ut.Translator

func init() {
	InitTans("zh")
}

func InitTans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			fieldName := fld.Name
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			if name == "" {
				//没有label就用json
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				return ","
			}
			return fmt.Sprintf("%s,%s", fieldName, name)
		})

		zhT := zh.New() //中文翻译
		enT := en.New() //英文翻译

		//第一个参数是备用(fallback)的语言环境
		//后面参数是应该支持的语言环境（支持多个）
		//uni := ut.New(zhT,zht)也是可以的
		uni := ut.New(enT, zhT, zhT)
		var ok bool
		//locale 通常取决与http请求头的'Accept-Language'
		//也可以使用 uni.FindTranslator
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		//注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用时候，需要传obj指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据错字段名，获取结构体的具体字段
			f, exits := getObj.Elem().FieldByName(e.Field())
			if exits {
				msg := f.Tag.Get("msg")
				if msg == "" {
					continue
				}
				return msg
			}
		}
	}
	return Error(err)
}

func Error(err error) (ret string) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	for _, e := range validationErrors {
		msg := e.Translate(Trans)
		oldFieldName := e.Field()
		_list := strings.Split(oldFieldName, ",")
		var fieldName string
		if len(_list) > 1 {
			fieldName = _list[0]
		}
		msg = strings.ReplaceAll(msg, fieldName, "")
		ret += msg + ";"
	}
	return ret
}

func ValidError(err error, obj any) (ret string, data map[string]string) {
	data = map[string]string{}
	getObj := reflect.TypeOf(obj)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error(), data
	}
	for _, e := range validationErrors {
		msg := e.Translate(Trans)
		oldFieldName := e.Field()
		_list := strings.Split(oldFieldName, ",")
		var fieldName string
		if len(_list) > 1 {
			fieldName = _list[0]
		}
		filed, ok := getObj.Elem().FieldByName(fieldName)
		if ok {
			msg = strings.ReplaceAll(msg, fieldName, "")
			jsonLabel, jsonOk := filed.Tag.Lookup("json")
			if jsonOk {
				data[jsonLabel] = msg
			}
		}
		ret += msg + ";"
	}
	return ret, data
}
