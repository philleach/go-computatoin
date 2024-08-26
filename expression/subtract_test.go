package expression

import (
	"testing"
)

func TestSubtractIsReducable(t *testing.T) {
	n := NewSubtract(NewNumber(3), NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestSubtractReduce(t *testing.T) {
	env := new(Environment)
	a := NewSubtract(NewNumber(3), NewNumber(2))
	res := a.Reduce(env)
	if res.Value() != 1 {
		t.Errorf("Reduce Returned %s, rather than (1)", res)
	}
}

func TestSubtractValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewSubtract(NewNumber(3), NewNumber(3))
	a.Value()
	t.Errorf("did not panic")
}
