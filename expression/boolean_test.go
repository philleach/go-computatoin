package expression

import (
	"testing"
)

func TestBoolean(t *testing.T) {
	b := Boolean(true)

	if !b.Value().(bool) {
		t.Errorf("Boolean(true) = %t; want true", b)
	}
}

func TestBooleanIsReducable(t *testing.T) {
	n := Boolean(true)

	if n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestBooleanReduce(t *testing.T) {
	defer func() { _ = recover() }()

	env := NewEnvironment()
	n := Boolean(false)
	n.Reduce(env)
	t.Errorf("did not panic")
}
