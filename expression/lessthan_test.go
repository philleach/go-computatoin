package expression

import (
	"testing"
)

func TestLessThanIsReducable(t *testing.T) {
	l := NewLessThan(NewNumber(3), NewNumber(3))

	if !l.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestLessThanReduce(t *testing.T) {
	env := new(Environment)
	l := NewLessThan(NewNumber(2), NewNumber(3))
	res := l.Reduce(env)
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestLessThanValue(t *testing.T) {
	defer func() { _ = recover() }()

	l := NewLessThan(NewNumber(3), NewNumber(3))
	l.Value()
	t.Errorf("did not panic")
}
