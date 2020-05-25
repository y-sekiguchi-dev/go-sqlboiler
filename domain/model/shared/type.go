package shared

type Type interface {
	Code() uint
	Symbol() string
}

type TypeImpl struct{
	code uint
	symbol string
}

func (t *TypeImpl) Code() uint {
	return t.code
}

func (t *TypeImpl) Symbol() string {
	return t.symbol
}

func NewTypeImpl(code uint, symbol string) *TypeImpl {
	return &TypeImpl{code, symbol}
}
