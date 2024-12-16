package syntax

import (
	"testing"
)

func TestIfIsReducable(t *testing.T) {
	n := If(Boolean(true), Assign("x", Number(3)), Assign("x", Number(3)))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return true")
	}
}

func TestIfReduce(t *testing.T) {
	env := NewEnvironment()
	a := If(Boolean(true), Assign("x", Number(3)), Assign("x", Number(1)))
	res, _ := a.Reduce(env)

	res2, env := res.Reduce(env)

	if !res2.IsReducable() && !(env.lookup("x").Value().(int) == 3) {
		t.Errorf("Reduce Returned %d, rather than (3)", env.lookup("x").Value().(int))
	}
	b := If(Boolean(false), Assign("x", Number(3)), Assign("x", Number(1)))
	res3, _ := b.Reduce(env)

	res4, env := res3.Reduce(env)

	if !res4.IsReducable() && !(env.lookup("x").Value().(int) == 1) {
		t.Errorf("Reduce Returned %d, rather than (1)", env.lookup("x").Value().(int))
	}
}
