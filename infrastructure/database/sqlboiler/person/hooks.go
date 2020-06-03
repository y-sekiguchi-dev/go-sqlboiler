package person

import (
	"context"
	"errors"
	"go-sqlboiler/infrastructure/database/sqlboiler/models"

	"github.com/volatiletech/sqlboiler/boil"
)

func ResistorHooks() {
	models.AddPersonHook(boil.BeforeInsertHook, setUserId)
	models.AddPersonHook(boil.BeforeUpdateHook, setUserId)
	models.AddPersonHook(boil.AfterSelectHook, loadChildren)
}

func setUserId(ctx context.Context, exec boil.ContextExecutor, p *models.Person) error {
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
	if err := p.Reload(ctx, exec); err != nil {
		return err
	}
	return loadChildren(ctx, exec, p)
}

func loadChildren(ctx context.Context, exec boil.ContextExecutor, p *models.Person) error {
	return p.L.LoadChildren(ctx, exec, true, p, nil)
}
