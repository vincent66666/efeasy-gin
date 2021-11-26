package initialize

import (
	"efeasy-gin/global"
	"efeasy-gin/utils"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func Validator(locale string) error {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.App.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.App.Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.App.Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, global.App.Trans)
		}

		// 注册自定义验证器
		RegisterValidatorFunc(v, "mobile", "手机号码非法", utils.ValidateMobile)
		return nil
	}
	return nil
}

// Func validator.ValidateMobile
type Func func(fl validator.FieldLevel) bool
// RegisterValidatorFunc 注册自定义校验tag
func RegisterValidatorFunc(v *validator.Validate, tag string, msgStr string, fn Func) {
	// 注册tag自定义校验
	_ = v.RegisterValidation(tag, validator.Func(fn))
	//自定义错误内容
	_ = v.RegisterTranslation(tag, global.App.Trans, func(ut ut.Translator) error {
		return ut.Add(tag, "{0}"+msgStr, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field())
		return t
	})

	return
}

