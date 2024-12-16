package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoNothingIsReducable(t *testing.T) {
	n := DoNothing()

	if n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestDoNothingEquality(t *testing.T) {
	n := DoNothing()
	m := DoNothing()

	if !cmp.Equal(n, m) {
		t.Errorf("DoNothing Equality failed")
	}
}
func TestDoNothingReduce(t *testing.T) {
	defer func() { _ = recover() }()

	env := NewEnvironment()
	n := DoNothing()
	n.Reduce(env)
	t.Errorf("did not panic")
}
