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

type provider struct {
	ctx       context.Context
	decorator func(fn interface{}) function.AnyFunc
}

func NewProvidedResult(ctx context.Context, decorator func(fn interface{}) function.AnyFunc) ProvidedResult {
	return &provider{
		ctx:       ctx,
		decorator: decorator,
	}
}

func (pr *provider) Context() context.Context {
	return pr.ctx
}

func (pr *provider) Transactional(fn interface{}) function.AnyFunc {
	return pr.decorator(fn)
}
