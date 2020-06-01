package transactional

import (
	"context"
	"database/sql/driver"
	"go-sqlboiler/infrastructure/reflect/function"
)

func decorate(target function.AnyFunc, tx driver.Tx, txCtx context.Context) function.AnyFunc {
	return func(args ...interface{}) function.Returns {
		returns := target(args...)
		if err := returns.Error(); err != nil {
			tx.Rollback()
			return function.ErrReturns(err)
		}
		if err := tx.Commit(); err != nil {
			return function.ErrReturns(err)
		}
		return returns
	}
}

// Param should be a function updates or inserts records across multiple tables.
// Param should return an error.
// It panics if fn breaks these preconditions.
func Function(fn interface{}, tx driver.Tx, ctx context.Context) function.AnyFunc {
	parsed := function.Parse(fn)
	return decorate(parsed, tx, ctx)
}