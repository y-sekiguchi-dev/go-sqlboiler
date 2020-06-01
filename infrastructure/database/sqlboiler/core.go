package sqlboiler

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"go-sqlboiler/infrastructure/database/transaction"
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

func ToExec(ctx context.Context) boil.ContextExecutor {
	if result, ok := ctx.(Context); ok {
		return result
	}
	return &readOnlyImpl{ctx,boil.GetContextDB()}
}

func newTransactionAwareContext(ctx context.Context) (*transactionAwareImpl, error) {
	if tx, err := boil.BeginTx(ctx, nil); err != nil {
		return nil, err
	} else {
		return &transactionAwareImpl{ctx,tx}, nil
	}
}

type providerImpl struct{}

func newTransactionProvider() transaction.Provider {
	return &providerImpl{}
}

func (i *providerImpl) Provide(ctx context.Context) (transaction.ProvidedResult, error) {
	if tx, err := newTransactionAwareContext(ctx); err != nil {
		return nil, err
	} else {
		decorator := func(fn interface{}) function.AnyFunc {
			return transaction.Decorate(fn, tx)
		}
		return transaction.NewProvidedResult(tx, decorator), nil
	}
}
