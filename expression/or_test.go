package expression

import (
	"testing"
)

func TestOrIsReducable(t *testing.T) {
	n := Or(Boolean(true), Boolean(true))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestOrReduce(t *testing.T) {
	env := NewEnvironment()
	n := Or(Boolean(true), Boolean(true))
	res := n.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestOrValue(t *testing.T) {
	defer func() { _ = recover() }()

	n := Or(Boolean(true), Boolean(true))
	n.Value()
	t.Errorf("did not panic")
}
