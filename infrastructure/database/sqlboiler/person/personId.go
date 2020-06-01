package person

import (
	"go-sqlboiler/domain/model"
	"strconv"
)

type idImpl struct {
	id uint
}

func (pid *idImpl) String() string {
	return strconv.Itoa(int(pid.id))
}

func (pid *idImpl) AsPersistForm() uint {
	return pid.id
}

func newPersonId(id uint) model.PersonId {
	return &idImpl{id}
}
