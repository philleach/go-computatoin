package expression

import (
	"testing"
)

func TestNewNumber(t *testing.T) {
	n := NewNumber(6)

	if n.Value() != 6 {
		t.Errorf("NewNumber(6) = %v; want 6", n)
	}
}

func TestIsReducable(t *testing.T) {
	n := NewNumber(6)

	if n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestReduce(t *testing.T) {
	defer func() { _ = recover() }()

	n := NewNumber(6)
	n.Reduce()
	t.Errorf("did not panic")
}
