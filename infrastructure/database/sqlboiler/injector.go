package sqlboiler

import (
	"context"
	"go-sqlboiler/domain/model"
	"go-sqlboiler/infrastructure/database/sqlboiler/person"
	"go-sqlboiler/infrastructure/reflect/function"
)

type Injector struct {
	ctx Context
}

func NewReadOnlyInjector(ctx context.Context) *Injector {
	return &Injector{newReadOnlyContext(ctx)}
}

func NewTransactionAwareInjector(ctx context.Context) (*Injector, error) {
	if ctx2, err := newTransactionAwareContext(ctx); err != nil {
		return nil, err
	} else {
		return &Injector{ctx2}, err
	}
}

func (i *Injector) NewPersonRepository() model.PersonRepository {
	return person.NewRepository(i.ctx)
}

func (i *Injector) Transactional(fn interface{}) function.AnyFunc {
	if tai, ok := i.ctx.(transactionAwareImpl); ok {
		return tai.transactional(fn)
	}
	panic("This injector instance is not transaction aware.")
}