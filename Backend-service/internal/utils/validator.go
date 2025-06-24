package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("phone", validatePhone)
}

func validatePhone(fl validator.FieldLevel) bool {
	phoneRegex := `^\+?[1-9]\d{1,14}$`
	return regexp.MustCompile(phoneRegex).MatchString(fl.Field().String())
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
