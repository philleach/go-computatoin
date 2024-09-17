package syntax

import (
	"testing"
)

func TestAddIsReducable(t *testing.T) {
	n := Add(Number(3), Number(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return true")
	}
}

func TestAddReduce(t *testing.T) {
	env := NewEnvironment()
	a := Add(Number(3), Number(3))
	res := a.Reduce(env)
	if res.Value() != 6 {
		t.Errorf("Reduce Returned %s, rather than (6)", res)
	}
}

func TestAddValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := Add(Number(3), Number(3))
	a.Value()
	t.Errorf("did not panic")
}
