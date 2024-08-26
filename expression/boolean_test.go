package expression

import (
	"testing"
)

func TestNewBoolean(t *testing.T) {
	b := NewBoolean(true)

	if !b.Value() {
		t.Errorf("NewBoolean(true) = %t; want true", b)
	}
}

func TestBooleanIsReducable(t *testing.T) {
	n := NewBoolean(true)

	if n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestBooleanReduce(t *testing.T) {
	defer func() { _ = recover() }()

	env := new(Environment)
	n := NewBoolean(false)
	n.Reduce(env)
	t.Errorf("did not panic")
}
