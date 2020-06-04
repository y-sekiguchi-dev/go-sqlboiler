package sqlboiler

import "testing"

func TestInit(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Error(err)
		}
	}()
	Init()()
}
