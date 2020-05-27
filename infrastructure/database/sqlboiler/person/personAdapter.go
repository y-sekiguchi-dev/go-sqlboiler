package person

import (
	"go-sqlboiler/domain/model"
	"go-sqlboiler/infrastructure/database/sqlboiler/models"
	"time"
)

func toDownStream(upstream model.Person) *models.Person {
	fullName := upstream.FullName()
	return &models.Person{
		PersonID:      0,
		FirstName:     fullName.FirstName(),
		LastName:      fullName.LastName(),
		Birthday:      time.Time{},
		Personality:   upstream.Personality(),
		HasPartner:    upstream.HasPartner(),
		Version:       int16(upstream.GetVersion()),
		Deleted:       upstream.Deleted(),
		CreatedUserID: 0,
		UpdatedUserID: 0,
	}
}

func toUpStream(downstream *models.Person) *model.Person {
	builder := model.AsStored(
			uint(downstream.PersonID),
			uint(downstream.Version),
			downstream.Deleted,
			model.NewBirthdayFromTime(downstream.Birthday),
			model.NewFullName(downstream.FirstName, downstream.LastName),
		)
	builder.Personality(model.GetPersonalityFrom(downstream.Personality))
	downstream.L.LoadChildren()
	return
}