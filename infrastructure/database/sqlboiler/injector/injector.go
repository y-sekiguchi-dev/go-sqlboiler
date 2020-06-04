package injector

import (
	person2 "go-sqlboiler/domain/model/person"
	"go-sqlboiler/infrastructure/database/sqlboiler"
	"go-sqlboiler/infrastructure/database/sqlboiler/person"
	"go-sqlboiler/infrastructure/database/transaction"
)

type Injector struct {
	ctx sqlboiler.Context
}

func (i *Injector) NewPersonRepository() person2.Repository {
	return person.NewRepository()
}

func (i *Injector) TransactionProvider() transaction.Provider {
	return sqlboiler.NewTransactionProvider()
}
