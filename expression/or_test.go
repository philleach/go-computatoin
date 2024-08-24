package expression

import (
	"testing"
)

func TestOrIsReducable(t *testing.T) {
	n := NewOr(NewBoolean(true), NewBoolean(true))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestOrReduce(t *testing.T) {
	n := NewOr(NewBoolean(true), NewBoolean(true))
	res := n.Reduce()
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestOrValue(t *testing.T) {
	defer func() { _ = recover() }()

	n := NewOr(NewBoolean(true), NewBoolean(true))
	n.Value()
	t.Errorf("did not panic")
}
