package expression

import (
	"fmt"
	"testing"
)

func TestAddIsReducable(t *testing.T) {
	n := NewAdd(NewNumber(3), NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestAddReduce(t *testing.T) {
	a := NewAdd(NewNumber(3), NewNumber(3))
	res := a.Reduce()
	fmt.Printf("Result = %s\n", res)
	if res != NewNumber(6) {
		t.Errorf("Reduce Returned %s, rather than (6)", res)
	}
}

func TestAddValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewAdd(NewNumber(3), NewNumber(3))
	a.Reduce()
	t.Errorf("did not panic")
}
