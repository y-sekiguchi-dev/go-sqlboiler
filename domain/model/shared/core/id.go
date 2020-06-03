package core

type Id interface {
	String() string
}

type emptyId struct{}

func (e *emptyId) String() string {
	return ""
}

var emptyIdInstance Id = &emptyId{}

func getEmptyId() Id {
	return emptyIdInstance
}

func IsEmpty(id Id) bool {
	return id == emptyIdInstance
}
