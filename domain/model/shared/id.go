package shared

type Id interface {
	String() string
}

type (
	emptyId Id
	emptyIdImpl struct {}
)

func (e *emptyIdImpl) String() string {
	return ""
}

var emptyIdInstance = emptyIdImpl{}

func getEmptyId() Id {
	return &emptyIdInstance
}

func IsEmpty(id Id) bool {
	_, ok := id.(emptyId)
	return ok
}