package model

type PersonRepository interface {
	FindById(id PersonId) (Person, error)
	Store(person Person) error
}

