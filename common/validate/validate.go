package validate


import (
	// "reflect"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Validate/v10 全局验证器
var trans ut.Translator

// var Validate *validator.Validate

func init() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)
		trans, _ = uni.GetTranslator("zh")

		var o bool
		local := "zh"
		trans, o = uni.GetTranslator(local)
		if !o {
			panic("uni.GetTranslator failed")
		}

		// 注册翻译器
		var err error
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		}

		if err != nil {
			panic("初始化验证语言失败" + err.Error())
		}

		v.RegisterValidation("checkNickname", checkNickname)
		v.RegisterTranslation("checkNickname", trans, func(ut ut.Translator) error {
			return ut.Add("checkNickname", "{0}长度必须大于4个字符!", false)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			s, _ := ut.T("checkNickname", fe.Field())
			return s
		})
		// 如果要启用请放开 import reflect
		/* v.RegisterTagNameFunc(func(field reflect.StructField) string {
			label := field.Tag.Get("label")
			if label == "" {
				return field.Name
			}
			return label
		}) */
	}

}

// 检验并返回检验错误信息
func Translate(err error) (errMsg string) {
	errs := err.(validator.ValidationErrors)
	for _, err := range errs {
		errMsg = err.Translate(trans)
		break
	}
	return
}

var checkNickname validator.Func = func(fl validator.FieldLevel) bool {
	if value, ok := fl.Field().Interface().(string); ok {
		// 长度大于6位
		return len(value) > 4
	}
	return true
}
