package expression

import (
	"testing"
)

func TestNumber(t *testing.T) {
	n := Number(6)

	if n.Value() != 6 {
		t.Errorf("Number(6) = %v; want 6", n)
	}
}

func TestNumberIsReducable(t *testing.T) {
	n := Number(6)

	if n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestNumberReduce(t *testing.T) {
	defer func() { _ = recover() }()

	env := NewEnvironment()
	n := Number(6)
	n.Reduce(env)
	t.Errorf("did not panic")
}
