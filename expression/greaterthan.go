package expression

import (
	"fmt"
)

type GreaterThanExpr struct {
	left  Expression
	right Expression
}

func GreaterThan(left Expression, right Expression) *GreaterThanExpr {
	g := new(GreaterThanExpr)
	g.left = left
	g.right = right
	return g
}

func (g GreaterThanExpr) String() string {
	return fmt.Sprintf("(%s > %s)", g.left, g.right)
}

func (g GreaterThanExpr) IsReducable() bool {
	return true
}

func (g GreaterThanExpr) Reduce(env *Environment) Expression {
	if g.left.IsReducable() {
		return GreaterThan(g.left.Reduce(env), g.right)
	} else if g.right.IsReducable() {
		return GreaterThan(g.left, g.right.Reduce(env))
	} else {
		left, ok := g.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := g.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Boolean(left > right)
		return res
	}
}

func (g GreaterThanExpr) Value() any {
	panic("Attempt to get value from GreaterThan")
}
