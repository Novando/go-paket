package contexts

import (
	"context"
)

func InjectCtx[T any](ctx context.Context, key string, val T) context.Context {
	return context.WithValue(ctx, key, val)
}
