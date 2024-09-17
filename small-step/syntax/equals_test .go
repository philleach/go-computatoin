package syntax

import (
	"testing"
)

func TestEqualsIsReducable(t *testing.T) {
	e := Equals(Number(3), Number(3))

	if !e.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestEqualsReduce(t *testing.T) {
	env := NewEnvironment()
	a := Equals(Number(3), Number(3))
	res := a.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestEqualsValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := Equals(Number(3), Number(3))
	a.Value()
	t.Errorf("did not panic")
}
