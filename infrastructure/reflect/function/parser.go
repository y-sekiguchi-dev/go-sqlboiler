package function

import "reflect"

// When it has an error, error() returns error and values has only one element of error.
type Returns interface {
	Error() error
	Value(index int) interface{}
}

type ReturnsImpl struct {
	values []interface{}
	errIndex int
}

func ErrReturns(err error) Returns {
	return &ReturnsImpl {
		[]interface{}{err},
		0,
	}
}

func (r *ReturnsImpl) Error() error {
	return r.values[r.errIndex].(error)
}

// Check error() returns nil before invoke.
// When wrapped function returns (int, string, error), value(0) returns int value,
// value(1) returns string and value(2) returns nil as type interface{}.
// Casting the return value to the original type will always success.
func (r *ReturnsImpl) Value(index int) interface{} {
	return r.values[index]
}

type AnyFunc func(...interface{}) Returns

func Parse(fn interface{}) AnyFunc {
	t := reflect.TypeOf(fn)
	if t.Kind() != reflect.Func {
		panic("param should be a function")
	}
	errType := reflect.TypeOf((*error)(nil)).Elem()
	errIdx := -1
	for i:=0; i<t.NumOut(); i++ {
		if t.Out(i).Implements(errType) {
			errIdx = i
			break
		}
	}
	if errIdx == -1 {
		panic("function should return an error.")
	}
	return func(args...interface{}) Returns {
		in := make([]reflect.Value, t.NumIn())
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}
		outputs := reflect.ValueOf(fn).Call(in)
		returns := make([]interface{}, t.NumOut())
		for i, output := range outputs {
			returns[i] = output.Interface()
		}
		return &ReturnsImpl{returns, errIdx}
	}
}
