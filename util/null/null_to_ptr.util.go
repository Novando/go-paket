package null

import (
	nullGuregu "github.com/guregu/null/v6"
)

func NullStringToPtr(v nullGuregu.String) *string {
	if !v.Valid {
		return nil
	}
	return &v.String
}
