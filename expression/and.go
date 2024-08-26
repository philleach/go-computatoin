package expression

import (
	"fmt"
)

type And struct {
	left  Expression[bool]
	right Expression[bool]
}

func NewAnd(left Expression[bool], right Expression[bool]) *And {
	a := new(And)
	a.left = left
	a.right = right
	return a
}

func (a And) String() string {
	return fmt.Sprintf("(%s && %s)", a.left, a.right)
}

func (a And) IsReducable() bool {
	return true
}

func (a And) Reduce(env *Environment) Expression[bool] {
	if a.left.IsReducable() {
		return NewAnd(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return NewAnd(a.left, a.right.Reduce(env))
	} else {
		res := NewBoolean(a.left.Value() && a.right.Value())
		return res
	}
}

func (a And) Value() bool {
	panic("Attempt to get value from And")
}
