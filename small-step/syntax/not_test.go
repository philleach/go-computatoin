package syntax

import (
	"testing"
)

func TestNotIsReducable(t *testing.T) {
	n := Not(Boolean(true))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return true")
	}
}

func TestNotReduce(t *testing.T) {
	env := NewEnvironment()
	n := Not(Boolean(false))
	res := n.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestNotValue(t *testing.T) {
	defer func() { _ = recover() }()

	n := Not(Boolean(true))
	n.Value()
	t.Errorf("did not panic")
}
