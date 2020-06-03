package core

type Entity interface {
	Id() Id
	Equals(entity Entity) bool
}

type entity struct {
	id Id
}

func IdenticalEntity(id Id) Entity {
	result := entity{id}
	return &result
}

func NewEntity() Entity {
	return IdenticalEntity(getEmptyId())
}

func (e *entity) Id() Id {
	return e.id
}

func (e *entity) Equals(entity Entity) bool {
	if IsEmpty(e.id) {
		return false
	}
	return e.id == entity.Id()
}
