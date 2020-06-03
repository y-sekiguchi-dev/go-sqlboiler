package sqlboiler

import (
	"context"
	"database/sql"
	"go-sqlboiler/infrastructure/database/transaction"
	"go-sqlboiler/infrastructure/reflect/function"

	"github.com/volatiletech/sqlboiler/boil"
)

type Context interface {
	context.Context
	boil.ContextExecutor
}

type readOnly struct {
	context.Context
	boil.ContextExecutor
}

type transactionAware struct {
	context.Context
	*sql.Tx
}

func ToExec(ctx context.Context) boil.ContextExecutor {
	if result, ok := ctx.(Context); ok {
		return result
	}
	return &readOnly{ctx, boil.GetContextDB()}
}

func newTransactionAwareContext(ctx context.Context) (*transactionAware, error) {
	if tx, err := boil.BeginTx(ctx, nil); err != nil {
		return nil, err
	} else {
		return &transactionAware{ctx, tx}, nil
	}
}

type provider struct{}

func newTransactionProvider() transaction.Provider {
	return &provider{}
}

func (i *provider) Provide(ctx context.Context) (transaction.ProvidedResult, error) {
	if tx, err := newTransactionAwareContext(ctx); err != nil {
		return nil, err
	} else {
		decorator := func(fn interface{}) function.AnyFunc {
			return transaction.Decorate(fn, tx)
		}
		return transaction.NewProvidedResult(tx, decorator), nil
	}
}
