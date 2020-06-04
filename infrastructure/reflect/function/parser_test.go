package function

import (
	"errors"
	"testing"
)

func TestErrReturns(t *testing.T) {
	// if ErrReturns called with an error
	err := errors.New("error")
	actual := ErrReturns(err)
	// then Error() should return the same error
	if actual.Error() != err {
		t.Errorf("%s should be %s", actual, err)
	}
}

func TestReturns_Error(t *testing.T) {
	// if returns has 3 values and errIndex 2
	expected := errors.New("c")
	target := returns{
		values:   []interface{}{"a", "b", expected},
		errIndex: 2,
	}
	// then returns.Error() should returns the value at index 2 of values
	actual := target.Error()
	if actual != expected {
		t.Errorf("%s should be %s", actual, expected)
	}
}

func TestParse(t *testing.T) {
	// if input function has 2 args and 3 returns whose tail is an error
	// if parsed function does not have an error
	inFunc := func(a int, b string) (int, string, error) {
		return a, b, nil
	}
	target := Parse(inFunc)(1, "a")
	// then Error() returns nil
	actual := target.Error()
	if target.Error() != nil {
		t.Errorf("%s should be %s", actual, "nil")
	}
	// then Value(i) returns (i)th output.
	expectedReturns := []interface{}{1, "a", nil}
	for i, expected := range expectedReturns {
		actual := target.Value(i)
		if actual != expected {
			t.Errorf("%s should be %s", actual, expected)
		}
	}

	// if parsed function has an error
	err := errors.New("error")
	inFunc = func(a int, b string) (int, string, error) {
		return 0, "", err
	}
	target = Parse(inFunc)(1, "a")
	// then Error() returns the error
	actual = target.Error()
	if target.Error() != err {
		t.Errorf("%s should be %s", actual, err)
	}
}

func TestParseNotAFunction(t *testing.T) {
	// if input is not a function
	// then panics
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("it should cause panic")
		}
	}()
	Parse("")
}

func TestParseNoErrReturningFunction(t *testing.T) {
	// if input function has 2 args and 2 returns without an error
	inFunc := func(a int, b string) (int, string) {
		return a, b
	}
	// then panics
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("it should cause panic")
		}
	}()
	Parse(inFunc)
}
