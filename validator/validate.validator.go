package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

func Validate(data interface{}) error {
	v := validator.New()
	_ = v.RegisterValidation("strong_password", strongPassword)
	_ = v.RegisterValidation("not_blank", notBlank)

	// Guregu v6
	v.RegisterCustomTypeFunc(gureguNullBoolV6, null.Bool{})
	v.RegisterCustomTypeFunc(gureguNullIntV6, null.Int{})
	v.RegisterCustomTypeFunc(gureguNullFloatV6, null.Float{})
	v.RegisterCustomTypeFunc(gureguNullStringV6, null.String{})
	v.RegisterCustomTypeFunc(gureguNullTimeV6, null.Time{})

	// Google UUID
	v.RegisterCustomTypeFunc(googleNullUUID, uuid.NullUUID{})
	v.RegisterCustomTypeFunc(googleUUID, uuid.UUID{})

	return v.Struct(data)
}
