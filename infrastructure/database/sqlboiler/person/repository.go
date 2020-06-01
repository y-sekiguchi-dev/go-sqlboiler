package person

import (
	"context"
	"errors"
	"github.com/volatiletech/sqlboiler/boil"
	domain "go-sqlboiler/domain/model"
	"go-sqlboiler/domain/model/shared"
	"go-sqlboiler/infrastructure/database/sqlboiler"
	"go-sqlboiler/infrastructure/database/sqlboiler/models"
)

type repository struct {
	adapter adapter
}

func NewRepository() domain.PersonRepository {
	return &repository{
		adapter: adapter{},
	}
}

func ResistorHooks() {
	models.AddPersonHook(boil.BeforeInsertHook, userIdSettingHook)
	models.AddPersonHook(boil.BeforeUpdateHook, userIdSettingHook)
	models.AddPersonHook(boil.AfterInsertHook, reload)
	models.AddPersonHook(boil.AfterUpdateHook, reload)
	models.AddPersonHook(boil.AfterSelectHook, loadChildren)
}

func userIdSettingHook(ctx context.Context, exec boil.ContextExecutor, p *models.Person) error {
	val := ctx.Value("userId")
	if val == nil {
		return errors.New("context does not have user id")
	}
	if userId, ok := val.(int64); ok {
		p.CreatedUserID = userId
		p.UpdatedUserID = userId
		return nil
	} else {
		return errors.New("user id should be an int value")
	}
}

func reload(ctx context.Context, exec boil.ContextExecutor, p *models.Person) error {
	if err:= p.Reload(ctx, exec); err !=nil {
		return err
	}
	return loadChildren(ctx, exec, p)
}

func loadChildren(ctx context.Context, exec boil.ContextExecutor, p *models.Person) error {
	return p.L.LoadChildren(ctx, exec, true, p, nil)
}

func (repo *repository) FindById(ctx context.Context, id domain.PersonId) (domain.Person, error) {
	if downstream, err := models.FindPerson(ctx, sqlboiler.ToExec(ctx), int64(id.AsPersistForm())); err != nil{
		return nil, err
	} else {
		return repo.adapter.toEntity(downstream), nil
	}
}

func (repo *repository) Store(ctx context.Context, person domain.Person) error {
	ds := repo.adapter.toDownStream(person)
	if shared.IsEmpty(person.Id()) {
		if err := ds.Insert(ctx, sqlboiler.ToExec(ctx), boil.Whitelist(
			"first_name", "last_name", "birthday", "personality", "create_user_id", "update_user_id",
			)); err != nil {
			return err
		}
		// TODO reset reloaded values like id.
		return nil
	} else {
		// TODO update
		return nil
	}

}
