package shared

type Entity interface {
	Id() Id
	Equals(entity Entity) bool
}

type EntityImpl struct {
	id Id
}

func IdenticalEntityImpl(id Id) *EntityImpl {
	result := EntityImpl{id}
	return &result
}

func NewEntityImpl() *EntityImpl {
	return IdenticalEntityImpl(getEmptyId())
}

func (e *EntityImpl) Id() Id {
	return e.id
}

func (e *EntityImpl) Equals(entity Entity) bool {
	if IsEmpty(e.id) {
		return false
	}
	return e.id == entity.Id()
}
