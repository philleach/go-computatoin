package expression

import (
	"testing"
)

func TestGreaterThanIsReducable(t *testing.T) {
	g := GreaterThan(Number(3), Number(3))

	if !g.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestGreaterThanReduce(t *testing.T) {
	env := NewEnvironment()
	g := GreaterThan(Number(4), Number(3))
	res := g.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestGreaterThanValue(t *testing.T) {
	defer func() { _ = recover() }()

	g := GreaterThan(Number(3), Number(3))
	g.Value()
	t.Errorf("did not panic")
}
