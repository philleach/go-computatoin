package expression

import (
	"testing"
)

func TestEqualsIsReducable(t *testing.T) {
	e := NewEquals(NewNumber(3), NewNumber(3))

	if !e.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestEqualsReduce(t *testing.T) {
	a := NewEquals(NewNumber(3), NewNumber(3))
	res := a.Reduce()
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestEqualsValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewEquals(NewNumber(3), NewNumber(3))
	a.Value()
	t.Errorf("did not panic")
}
