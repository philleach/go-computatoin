package syntax

import (
	"testing"
)

func TestMultiplyIsReducable(t *testing.T) {
	n := Multiply(Number(3), Number(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestMultiplyReduce(t *testing.T) {
	env := NewEnvironment()
	a := Multiply(Number(3), Number(3))
	res := a.Reduce(env)
	if res.Value() != 9 {
		t.Errorf("Reduce Returned %s, rather than (9)", res)
	}
}

func TestMultiplyValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := Multiply(Number(3), Number(3))
	a.Value()
	t.Errorf("did not panic")
}
