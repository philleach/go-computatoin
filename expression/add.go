package expression

import (
	"fmt"
)

type AddExpr struct {
	left  Expression
	right Expression
}

func Add(left Expression, right Expression) *AddExpr {
	a := new(AddExpr)
	a.left = left
	a.right = right
	return a
}

func (a AddExpr) String() string {
	return fmt.Sprintf("(%s + %s)", a.left, a.right)
}

func (a AddExpr) IsReducable() bool {
	return true
}

func (a AddExpr) Reduce(env *Environment) Expression {
	if a.left.IsReducable() {
		return Add(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return Add(a.left, a.right.Reduce(env))
	} else {
		left, ok := a.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := a.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Number(left + right)
		return res
	}
}

func (a AddExpr) Value() any {
	panic("Attempt to get value from Add")
}
