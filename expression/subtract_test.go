package expression

import (
	"testing"
)

func TestSubtractIsReducable(t *testing.T) {
	n := Subtract(Number(3), Number(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestSubtractReduce(t *testing.T) {
	env := NewEnvironment()
	a := Subtract(Number(3), Number(2))
	res := a.Reduce(env)
	if res.Value() != 1 {
		t.Errorf("Reduce Returned %s, rather than (1)", res)
	}
}

func TestSubtractValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := Subtract(Number(3), Number(3))
	a.Value()
	t.Errorf("did not panic")
}
