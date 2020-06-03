package core

type EnumType interface {
	Code() uint
	Symbol() string
}

type enumType struct {
	code   uint
	symbol string
}

func (t *enumType) Code() uint {
	return t.code
}

func (t *enumType) Symbol() string {
	return t.symbol
}

func NewType(code uint, symbol string) EnumType {
	return &enumType{code, symbol}
}
