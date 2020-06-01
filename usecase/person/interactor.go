package person

import (
	"context"
	"go-sqlboiler/domain/model"
	"go-sqlboiler/infrastructure/database/transaction"
)

type UseCase interface {
	Register(ctx context.Context, person model.Person) error
	Revise(ctx context.Context, person model.Person) error
	Find(ctx context.Context, id model.PersonId) (model.Person, error)
}

type impl struct {
	repo model.PersonRepository
	tx transaction.Provider
}

func NewUseCase(repo model.PersonRepository, tx transaction.Provider) UseCase {
	return &impl{repo: repo, tx: tx}
}

func (i *impl) Find(ctx context.Context, id model.PersonId) (model.Person, error) {
	return i.repo.FindById(ctx, id)
}

func (i *impl) Register(ctx context.Context, person model.Person) error {
	if pr, err := i.tx.Provide(ctx); err != nil {
		return err
	} else {
		return pr.Transactional(i.repo.Store)(pr.Context(), person).Error()
	}
}

func (i *impl) Revise(ctx context.Context, person model.Person) error {
	if pr, err := i.tx.Provide(ctx); err != nil {
		return err
	} else {
		return pr.Transactional(i.repo.Store)(pr.Context(), person).Error()
	}
}