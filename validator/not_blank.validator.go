package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func notBlank(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}
