package syntax

import (
	"testing"
)

func TestLessThanIsReducable(t *testing.T) {
	l := LessThan(Number(3), Number(3))

	if !l.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestLessThanReduce(t *testing.T) {
	env := NewEnvironment()
	l := LessThan(Number(2), Number(3))
	res := l.Reduce(env)
	if !res.Value().(bool) {
		t.Errorf("Reduce Returned %t, rather than (true)", res)
	}
}

func TestLessThanValue(t *testing.T) {
	defer func() { _ = recover() }()

	l := LessThan(Number(3), Number(3))
	l.Value()
	t.Errorf("did not panic")
}
