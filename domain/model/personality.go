package model

import "go-sqlboiler/domain/model/shared"

type Personality struct {
	*shared.TypeImpl
}

var (
	ANALYSTS = Personality{shared.NewTypeImpl(1, "ANALYSTS")}
	DIPLOMATS = Personality{shared.NewTypeImpl(2, "DIPLOMATS")}
	SENTINELS = Personality{shared.NewTypeImpl(3, "SENTINELS")}
	EXPLORES = Personality{shared.NewTypeImpl(4, "EXPLORES")}
)
