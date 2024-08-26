package expression

import (
	"testing"
)

func TestGreaterThanIsReducable(t *testing.T) {
	g := NewGreaterThan(NewNumber(3), NewNumber(3))

	if !g.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestGreaterThanReduce(t *testing.T) {
	env := new(Environment)
	g := NewGreaterThan(NewNumber(4), NewNumber(3))
	res := g.Reduce(env)
	if !res.Value() {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestGreaterThanValue(t *testing.T) {
	defer func() { _ = recover() }()

	g := NewGreaterThan(NewNumber(3), NewNumber(3))
	g.Value()
	t.Errorf("did not panic")
}
