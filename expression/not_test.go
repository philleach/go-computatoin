package expression

import (
	"testing"
)

func TestNotIsReducable(t *testing.T) {
	n := NewNot(NewBoolean(true))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return true")
	}
}

func TestNotReduce(t *testing.T) {
	env := new(Environment)
	n := NewNot(NewBoolean(false))
	res := n.Reduce(env)
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestNotValue(t *testing.T) {
	defer func() { _ = recover() }()

	n := NewNot(NewBoolean(true))
	n.Value()
	t.Errorf("did not panic")
}
