package person

import (
	"go-sqlboiler/domain/model/person"
	sqlboiler "go-sqlboiler/infrastructure/database/sqlboiler/models"
)

type childAdapter struct {
}

func (ca *childAdapter) toDownStream(upstream person.Child) *sqlboiler.Child {
	fullName := upstream.FullName()
	return &sqlboiler.Child{
		SubNo:     int8(upstream.SubNo()),
		FirstName: fullName.FirstName(),
		LastName:  fullName.LastName(),
		Birthday:  upstream.Birthday().AsTime(),
	}
}

func (ca *childAdapter) toDownStreams(upstreams []person.Child) []*sqlboiler.Child {
	result := make([]*sqlboiler.Child, 0, len(upstreams))
	for _, upstream := range upstreams {
		result = append(result, ca.toDownStream(upstream))
	}
	return result
}

func (ca *childAdapter) toEntity(downstream *sqlboiler.Child) person.Child {
	birthday, _ := person.NewBirthdayFromTime(downstream.Birthday)
	return person.StoredChild(
		uint(downstream.SubNo),
		person.NewFullName(downstream.FirstName, downstream.LastName),
		birthday,
	)
}
