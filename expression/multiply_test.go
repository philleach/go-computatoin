package expression

import (
	"testing"
)

func TestMultiplyIsReducable(t *testing.T) {
	n := NewMultiply(NewNumber(3), NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestMultiplyReduce(t *testing.T) {
	env := new(Environment)
	a := NewMultiply(NewNumber(3), NewNumber(3))
	res := a.Reduce(env)
	if res.Value() != 9 {
		t.Errorf("Reduce Returned %s, rather than (9)", res)
	}
}

func TestMultiplyValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewMultiply(NewNumber(3), NewNumber(3))
	a.Value()
	t.Errorf("did not panic")
}
