package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}

	// 用 label tag 作为字段显示名，没有 label 就用字段名
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		if label := field.Tag.Get("label"); label != "" {
			return label
		}
		return field.Name
	})

	_ = zhTranslations.RegisterDefaultTranslations(v, trans)
}

// ValidateErr 返回中文错误字符串
func ValidateErr(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	var msgs []string
	for _, e := range errs {
		msgs = append(msgs, e.Translate(trans))
	}
	return strings.Join(msgs, "; ")
}

// ValidateFieldErrors 返回字段级错误 map
func ValidateFieldErrors(err error) map[string]string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"_error": err.Error()}
	}

	result := make(map[string]string)
	for _, e := range errs {
		// e.Field() 现在是 label 名
		result[e.Field()] = e.Translate(trans)
	}
	return result
}
