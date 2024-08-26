package expression

import (
	"testing"
)

func TestDivideIsReducable(t *testing.T) {
	n := NewDivide(NewNumber(3), NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestDivideReduce(t *testing.T) {
	env := new(Environment)
	a := NewDivide(NewNumber(3), NewNumber(3))
	res := a.Reduce(env)
	if res.Value() != 1 {
		t.Errorf("Reduce Returned %s, rather than (1)", res)
	}
}

func TestDivideValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewDivide(NewNumber(3), NewNumber(3))
	a.Value()
	t.Errorf("did not panic")
}
