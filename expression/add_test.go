package expression

import (
	"testing"
)


func TestAddIsReducable(t *testing.T) {
	n := NewAdd(NewNumber(3),NewNumber(3))

	if !n.IsReducable() {
		t.Errorf("IsReducable should return false")
	}
}

func TestAddReduce(t *testing.T) {
func TestAddValue(t *testing.T) {
	defer func() { _ = recover() }()

	a := NewAdd(NewNumber(3),NewNumber(3))
	a.Reduce()
	t.Errorf("did not panic")
}
