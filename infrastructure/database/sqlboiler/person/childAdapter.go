package person

import (
	"go-sqlboiler/domain/model"
	"go-sqlboiler/infrastructure/database/sqlboiler/models"
	"time"
)

func toDownStream(personId model.PersonId, upstream *model.Child) *models.Child {
	fullName := upstream.FullName()
	personId.
	return &models.Child{
		PersonID:  0,
		SubNo:     int8(upstream.SubNo()),
		FirstName: fullName.FirstName(),
		LastName:  fullName.LastName(),
		Birthday:  time.Time{}
	}
}

func toUpStream(downstream *models.Person) *model.Person {