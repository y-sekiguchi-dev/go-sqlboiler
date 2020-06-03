package person

import "context"

type Repository interface {
	FindById(ctc context.Context, id Id) (Person, error)
	Store(ctc context.Context, person Person) error
}
