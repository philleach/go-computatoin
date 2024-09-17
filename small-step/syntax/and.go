package syntax

import (
	"fmt"
)

type AndExpr struct {
	left  Expression
	right Expression
}

func And(left Expression, right Expression) *AndExpr {
	a := new(AndExpr)
	a.left = left
	a.right = right
	return a
}

func (a AndExpr) String() string {
	return fmt.Sprintf("(%s && %s)", a.left, a.right)
}

func (a AndExpr) IsReducable() bool {
	return true
}

func (a AndExpr) Reduce(env *Environment) Expression {
	if a.left.IsReducable() {
		return And(a.left.Reduce(env), a.right)
	} else if a.right.IsReducable() {
		return And(a.left, a.right.Reduce(env))
	} else {
		left, ok := a.left.Value().(bool)
		if !ok {
			panic("Left value is not an bool")
		}
		right, ok := a.right.Value().(bool)
		if !ok {
			panic("Right value is not an bool")
		}
		res := Boolean(left && right)
		return res
	}
}

func (a AndExpr) Value() any {
	panic("Attempt to get value from And")
}
