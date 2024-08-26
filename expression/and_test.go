package expression

import (
	"testing"
)

func TestAndIsReducable(t *testing.T) {
	n := NewAnd(NewBoolean(true), NewBoolean(false))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestAndReduce(t *testing.T) {
	env := new(Environment)
	a := NewAnd(NewBoolean(true), NewBoolean(true))
	res := a.Reduce(env)
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestAndValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewAnd(NewBoolean(true), NewBoolean(false))
	a.Value()
	t.Errorf("did not panic")
}
