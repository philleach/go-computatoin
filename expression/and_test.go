package expression

import (
	"testing"
)

func TestAndIsReducable(t *testing.T) {
	n := And(Boolean(true), Boolean(false))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestAndReduce(t *testing.T) {
	env := NewEnvironment()
	a := And(Boolean(true), Boolean(true))
	res := a.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestAndValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := And(Boolean(true), Boolean(false))
	a.Value()
	t.Errorf("did not panic")
}
