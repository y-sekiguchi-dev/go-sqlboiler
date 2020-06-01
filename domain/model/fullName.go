package model

type FullName struct {
	firstName string
	lastName string
}

func NewFullName(firstName string, lastName string) FullName {
	return FullName{firstName, lastName}
}

func (f FullName) AsString(sep string) string {
	return f.firstName + sep + f.lastName
}

func (f FullName) WithFirstName(firstName string) FullName {
	return NewFullName(firstName, f.lastName)
}

func (f FullName) WithLastName(lastName string) FullName {
	return NewFullName(f.firstName, lastName)
}

func (f FullName) FirstName() string {
	return f.firstName
}

func (f FullName) LastName() string {
	return f.lastName
}