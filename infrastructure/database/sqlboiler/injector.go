package sqlboiler

import (
	"go-sqlboiler/domain/model"
	"go-sqlboiler/infrastructure/database/sqlboiler/person"
	"go-sqlboiler/infrastructure/database/transaction"
)

type Injector struct {
	ctx Context
}

func (i *Injector) NewPersonRepository() model.PersonRepository {
	return person.NewRepository()
}

func (i *Injector) TransactionProvider() transaction.Provider {
	return newTransactionProvider()
}