package person

import (
	"github.com/volatiletech/null"
	domain "go-sqlboiler/domain/model"
	sqlboiler "go-sqlboiler/infrastructure/database/sqlboiler/models"
	"time"
)

type adapter struct {

}

func (a *adapter) toDownStream(entity domain.Person) *sqlboiler.Person {
	fullName := entity.FullName()
	personality := null.Int8FromPtr(nil)
	if entity.Personality() != nil {
		personality = null.Int8From(int8(entity.Personality().Code()))
	}
	return &sqlboiler.Person{
		FirstName:   fullName.FirstName(),
		LastName:    fullName.LastName(),
		Birthday:    time.Time{},
		Personality: personality,
		HasPartner:  entity.HasPartner(),
		Version:     int16(entity.GetVersion()),
		Deleted:     entity.Deleted(),
	}
}

func (a *adapter) toEntity(downstream *sqlboiler.Person) domain.Person {
	birthday, _ := domain.NewBirthdayFromTime(downstream.Birthday)
	builder := domain.AsStored(
			newPersonId(uint(downstream.PersonID)),
			uint(downstream.Version),
			downstream.Deleted,
			birthday,
			domain.NewFullName(downstream.FirstName, downstream.LastName),
		)
	if downstream.Personality.Valid {
		builder.Personality(domain.GetPersonalityFrom(uint(downstream.Personality.Int8)))
	}
	ca := childAdapter{}
	for _, child := range downstream.R.Children {
		entity := ca.toEntity(child)
		builder.AddChild(entity)
	}
	return builder.Build()
}