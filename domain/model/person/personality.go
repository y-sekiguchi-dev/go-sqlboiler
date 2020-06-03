package person

import (
	"go-sqlboiler/domain/model/shared/core"
)

type Personality struct {
	core.EnumType
}

var (
	// TODO Defining NULL as TypeImpl{0, ""} can avoid nullable.
	ANALYSTS  = Personality{core.NewType(1, "ANALYSTS")}
	DIPLOMATS = Personality{core.NewType(2, "DIPLOMATS")}
	SENTINELS = Personality{core.NewType(3, "SENTINELS")}
	EXPLORES  = Personality{core.NewType(4, "EXPLORES")}
)

var codeTypeMap = map[uint]*Personality{
	ANALYSTS.Code():  &ANALYSTS,
	DIPLOMATS.Code(): &DIPLOMATS,
	SENTINELS.Code(): &SENTINELS,
	EXPLORES.Code():  &EXPLORES,
}

func GetPersonalityFrom(code uint) *Personality {
	return codeTypeMap[code]
}
