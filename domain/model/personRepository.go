package model

type PersonRepository interface {
	FindById(id PersonId) Person
	Store(person Person)
}

