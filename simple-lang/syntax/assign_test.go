package syntax

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAssignIsReducable(t *testing.T) {
	n := Assign("x", Number(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return true")
	}
}

func TestAssignReduce(t *testing.T) {
	env := NewEnvironment()
	a := Assign("x", Number(3))
	res, newEnv := a.Reduce(env)
	m := DoNothing()

	if !cmp.Equal(res, m) {
		t.Errorf("Reduce Returned %s, rather than (6)", res)
	}
	if newEnv.lookup("x").Value() != 3 {
		t.Error("Environment doesn't contain 3")
	}
}
