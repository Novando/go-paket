package validator

import (
	"reflect"

	"github.com/google/uuid"
)

func googleNullUUID(r reflect.Value) any {
	if v, ok := r.Interface().(uuid.NullUUID); ok {
		if v.Valid {
			return v.UUID.String()
		}
	}
	return nil
}

func googleUUID(r reflect.Value) any {
	if v, ok := r.Interface().(uuid.UUID); ok {
		return v.String()
	}
	return nil
}
