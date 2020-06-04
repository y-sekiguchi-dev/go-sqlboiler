package person

import (
	"context"
	"errors"
	"go-sqlboiler/domain/model/person"
	"go-sqlboiler/domain/model/shared/core"
	"go-sqlboiler/infrastructure/database/sqlboiler"
	"go-sqlboiler/infrastructure/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/boil"

	. "github.com/volatiletech/sqlboiler/queries/qm"
)

type repository struct {
	adapter
	childAdapter
}

func NewRepository() person.Repository {
	return &repository{
		adapter:      adapter{},
		childAdapter: childAdapter{},
	}
}

func (repo *repository) FindById(ctx context.Context, id person.Id) (person.Person, error) {
	if downstream, err := models.FindPerson(ctx, sqlboiler.ToExec(ctx), int64(id.AsPersistForm())); err != nil {
		return nil, err
	} else {
		return repo.adapter.toEntity(downstream), nil
	}
}

func (repo *repository) Store(ctx context.Context, person person.Person) error {
	if core.IsEmpty(person.Id()) {
		return repo.insert(ctx, person)
	}
	return repo.update(ctx, person)
}

func (repo *repository) insert(ctx context.Context, person person.Person) error {
	ds := repo.adapter.toDownStream(person)
	if err := ds.Insert(ctx, sqlboiler.ToExec(ctx), boil.Blacklist("person_id", "created_at", "updated_at")); err != nil {
		return err
	}
	if err := repo.storeChildren(ctx, ds, person.ChildrenView()); err != nil {
		return err
	}
	person.GiveId(newPersonId(uint(ds.PersonID)))
	return nil
}

func (repo *repository) update(ctx context.Context, person person.Person) error {
	ds := repo.adapter.toDownStream(person)
	cnt, err := models.Persons(Where("person_id = ?", ds.PersonID),
		And("version = ?", ds.Version)).UpdateAll(
		ctx, sqlboiler.ToExec(ctx), models.M{
			models.PersonColumns.Version:       ds.Version + 1,
			models.PersonColumns.UpdatedUserID: ds.UpdatedUserID,
			models.PersonColumns.CreatedUserID: ds.CreatedUserID,
			models.PersonColumns.Birthday:      ds.Birthday,
			models.PersonColumns.LastName:      ds.LastName,
			models.PersonColumns.FirstName:     ds.FirstName,
			models.PersonColumns.Personality:   ds.Personality,
			models.PersonColumns.HasPartner:    ds.HasPartner,
			models.PersonColumns.Deleted:       ds.Deleted,
		})
	if err != nil {
		return err
	}
	if cnt == 0 {
		return errors.New("optimistic lock was failed")
	}
	if err := repo.storeChildren(ctx, ds, person.ChildrenView()); err != nil {
		return err
	}
	person.IncrementVersion()
	return nil
}

func (repo *repository) storeChildren(ctx context.Context, ds *models.Person, children []person.Child) error {
	exec := sqlboiler.ToExec(ctx)
	// delete all
	if _, err := models.Children(models.ChildWhere.PersonID.EQ(ds.PersonID)).DeleteAll(ctx, exec); err != nil {
		return err
	}
	// insert all
	childrenD := repo.childAdapter.toDownStreams(children)
	return ds.AddChildren(ctx, sqlboiler.ToExec(ctx), true, childrenD...)
}
