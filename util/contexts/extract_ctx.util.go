package contexts

import (
	"context"
)

func ExtractCtx[T any](ctx context.Context, key string) (T, bool) {
	tx, ok := ctx.Value(key).(T)
	return tx, ok
}
