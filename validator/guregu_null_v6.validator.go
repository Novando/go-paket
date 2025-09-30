package validator

import (
	"reflect"

	"github.com/guregu/null/v6"
)

func gureguNullStringV6(r reflect.Value) any {
	if v, ok := r.Interface().(null.String); ok {
		if v.Valid {
			return v.String
		}
	}
	return nil
}

func gureguNullTimeV6(r reflect.Value) any {
	if v, ok := r.Interface().(null.Time); ok {
		if v.Valid {
			return v.Time
		}
	}
	return nil
}

func gureguNullFloatV6(r reflect.Value) any {
	if v, ok := r.Interface().(null.Float); ok {
		if v.Valid {
			return v.Float64
		}
	}
	return nil
}

func gureguNullIntV6(r reflect.Value) any {
	if v, ok := r.Interface().(null.Int); ok {
		if v.Valid {
			return v.Int64
		}
	}
	return nil
}

func gureguNullBoolV6(r reflect.Value) any {
	if v, ok := r.Interface().(null.Bool); ok {
		if v.Valid {
			return v.Bool
		}
	}
	return nil
}
