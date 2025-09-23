package pgx

import (
	"context"

	"github.com/Novando/go-paket/constant"
	"github.com/Novando/go-paket/util/contexts"
	p "github.com/jackc/pgx/v5"
)

func HandleTx(ctx context.Context, err error) {
	tx, ok := contexts.ExtractCtx[p.Tx](ctx, constant.ContextTxKey)
	if !ok {
		return
	}
	if err != nil {
		_ = tx.Rollback(ctx)
	} else {
		_ = tx.Commit(ctx)
	}
}
