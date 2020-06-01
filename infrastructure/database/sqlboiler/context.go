package sqlboiler

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"go-sqlboiler/infrastructure/database/transactional"
	"go-sqlboiler/infrastructure/reflect/function"
)

type Context interface {
	context.Context
	boil.ContextExecutor
}

type readOnlyImpl struct {
	context.Context
	boil.ContextExecutor
}

type transactionAwareImpl struct {
	context.Context
	*sql.Tx
}

func newReadOnlyContext(ctx context.Context) Context {
	return &readOnlyImpl{ctx,boil.GetContextDB()}
}

func newTransactionAwareContext(ctx context.Context) (Context, error) {
	if tx, err := boil.BeginTx(ctx, nil); err != nil {
		return nil, err
	} else {
		return &transactionAwareImpl{ctx,tx}, nil
	}
}

func (t *transactionAwareImpl) transactional(fn interface{}) function.AnyFunc {
	return transactional.Function(fn, t, t)
}
