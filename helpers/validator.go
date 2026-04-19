package helpers

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func Validate(model any) ([]string, bool) {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(model)

	var resultErrors []string
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			resultErrors = append(resultErrors, err.Translate(trans))
		}
	}

	if resultErrors != nil{
		return resultErrors, false
	}else{
		return nil, true
	}
}
