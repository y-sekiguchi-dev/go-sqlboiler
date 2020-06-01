package person

import (
	domain "go-sqlboiler/domain/model"
	sqlboiler "go-sqlboiler/infrastructure/database/sqlboiler/models"
)

type childAdapter struct {

}

func (ca *childAdapter) toDownStream(upstream domain.Child) *sqlboiler.Child {
	fullName := upstream.FullName()
	return &sqlboiler.Child{
		SubNo:     int8(upstream.SubNo()),
		FirstName: fullName.FirstName(),
		LastName:  fullName.LastName(),
		Birthday: upstream.Birthday().AsTime(),
	}
}

func (ca *childAdapter) toEntity(downstream *sqlboiler.Child) domain.Child {
	birthday, _ := domain.NewBirthdayFromTime(downstream.Birthday)
	return domain.NewChild(
		uint(downstream.SubNo),
		domain.NewFullName(downstream.FirstName, downstream.LastName),
		birthday,
		)
}