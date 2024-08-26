package expression

import (
	"testing"
)

func TestAddIsReducable(t *testing.T) {
	n := NewAdd(NewNumber(3), NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestAddReduce(t *testing.T) {
	env := new(Environment)
	a := NewAdd(NewNumber(3), NewNumber(3))
	res := a.Reduce(env)
	if res.Value() != 6 {
		t.Errorf("Reduce Returned %s, rather than (6)", res)
	}
}

func TestAddValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewAdd(NewNumber(3), NewNumber(3))
	a.Value()
	t.Errorf("did not panic")
}
