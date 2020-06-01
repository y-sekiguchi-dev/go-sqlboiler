package transaction

import (
	"context"
	"go-sqlboiler/infrastructure/reflect/function"
)

type Provider interface {
	Provide(ctx context.Context) (ProvidedResult, error)
}

type ProvidedResult interface {
	Context() context.Context
	Transactional(fn interface{}) function.AnyFunc
}

type impl struct {
	ctx context.Context
	decorator func(fn interface{}) function.AnyFunc
}

func NewProvidedResult(ctx context.Context, decorator func(fn interface{}) function.AnyFunc) ProvidedResult {
	return &impl{
		ctx:       ctx,
		decorator: decorator,
	}
}

func (pr *impl) Context() context.Context {
	return pr.ctx
}

func (pr *impl) Transactional(fn interface{}) function.AnyFunc {
	return pr.decorator(fn)
}