package model

import "go-sqlboiler/domain/model/shared"

type Personality struct {
	*shared.TypeImpl
}

var (
	// TODO Defining NULL as TypeImpl{0, ""} can avoid nullable.
	ANALYSTS = Personality{shared.NewTypeImpl(1, "ANALYSTS")}
	DIPLOMATS = Personality{shared.NewTypeImpl(2, "DIPLOMATS")}
	SENTINELS = Personality{shared.NewTypeImpl(3, "SENTINELS")}
	EXPLORES = Personality{shared.NewTypeImpl(4, "EXPLORES")}
)

var codeTypeMap = map[uint]*Personality{
	ANALYSTS.Code(): &ANALYSTS,
	DIPLOMATS.Code(): &DIPLOMATS,
	SENTINELS.Code(): &SENTINELS,
	EXPLORES.Code(): &EXPLORES,
}

func GetPersonalityFrom(code uint) *Personality {
	return codeTypeMap[code]
}