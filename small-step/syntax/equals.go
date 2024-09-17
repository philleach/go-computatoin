package syntax

import (
	"fmt"
)

type EqualsExpr struct {
	left  Expression
	right Expression
}

func Equals(left Expression, right Expression) *EqualsExpr {
	e := new(EqualsExpr)
	e.left = left
	e.right = right
	return e
}

func (e EqualsExpr) String() string {
	return fmt.Sprintf("%s = %s", e.left, e.right)
}

func (e EqualsExpr) IsReducable() bool {
	return true
}

func (e EqualsExpr) Reduce(env *Environment) Expression {
	if e.left.IsReducable() {
		return Equals(e.left.Reduce(env), e.right)
	} else if e.right.IsReducable() {
		return Equals(e.left, e.right.Reduce(env))
	} else {
		left, ok := e.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := e.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Boolean(left == right)
		return res
	}
}

func (e EqualsExpr) Value() any {
	panic("Attempt to get value from Equals")
}
