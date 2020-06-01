package model

import "context"

type PersonRepository interface {
	FindById(ctc context.Context, id PersonId) (Person, error)
	Store(ctc context.Context, person Person) error
}

