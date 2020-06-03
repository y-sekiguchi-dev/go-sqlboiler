package person

import (
	"context"
	"go-sqlboiler/domain/model/person"
	"go-sqlboiler/infrastructure/database/transaction"
)

type UseCase interface {
	Register(ctx context.Context, person person.Person) error
	Revise(ctx context.Context, person person.Person) error
	Find(ctx context.Context, id person.Id) (person.Person, error)
}

type useCase struct {
	repo person.Repository
	tx   transaction.Provider
}

func NewUseCase(repo person.Repository, tx transaction.Provider) UseCase {
	return &useCase{repo: repo, tx: tx}
}

func (i *useCase) Find(ctx context.Context, id person.Id) (person.Person, error) {
	return i.repo.FindById(ctx, id)
}

func (i *useCase) Register(ctx context.Context, person person.Person) error {
	if pr, err := i.tx.Provide(ctx); err != nil {
		return err
	} else {
		return pr.Transactional(i.repo.Store)(pr.Context(), person).Error()
	}
}

func (i *useCase) Revise(ctx context.Context, person person.Person) error {
	if pr, err := i.tx.Provide(ctx); err != nil {
		return err
	} else {
		return pr.Transactional(i.repo.Store)(pr.Context(), person).Error()
	}
}
