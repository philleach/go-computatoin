package expression

import (
	"fmt"
)

type Add struct {
	left  Expression[int]
	right Expression[int]
}

func NewAdd(left Expression[int], right Expression[int]) *Add {
	a := new(Add)
	a.left = left
	a.right = right
	return a
}

func (a Add) String() string {
	return fmt.Sprintf("(%s + %s)", a.left, a.right)
}

func (a Add) IsReducable() bool {
	return true
}

func (a Add) Reduce(env *Environment) Expression[int] {
	if a.left.IsReducable() {
		return NewAdd(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return NewAdd(a.left, a.right.Reduce(env))
	} else {
		res := NewNumber(a.left.Value() + a.right.Value())
		return res
	}
}

func (a Add) Value() int {
	panic("Attempt to get value from Add")
}
