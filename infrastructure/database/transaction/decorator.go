package transaction

import (
	"database/sql/driver"
	"go-sqlboiler/infrastructure/reflect/function"
)

func decorate(target function.AnyFunc, tx driver.Tx) function.AnyFunc {
	abort := func (err error) function.Returns {
		tx.Rollback()
		return function.ErrReturns(err)
	}
	return func(args ...interface{}) function.Returns {
		var result function.Returns
		defer func() {
			if p := recover(); p!=nil {
				tx.Rollback()
				panic(p)
			}
			if err := result.Error(); err != nil {
				result = abort(err)
			}
			if err := tx.Commit(); err != nil {
				result = abort(err)
			}
		}()
		result = target(args...)
		return result
	}
}

// Param should be a function updates or inserts records and needs to be atomic.
// Param should return an error.
// It panics if fn breaks these preconditions.
func Decorate(fn interface{}, tx driver.Tx) function.AnyFunc {
	parsed := function.Parse(fn)
	return decorate(parsed, tx)
}