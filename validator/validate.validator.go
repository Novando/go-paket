package validator

import (
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) error {
	validate := validator.New()
	_ = validate.RegisterValidation("strong_password", strongPassword)
	_ = validate.RegisterValidation("not_blank", notBlank)

	return validate.Struct(data)
}
