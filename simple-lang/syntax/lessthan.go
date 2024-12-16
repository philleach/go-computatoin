package syntax

import (
	"fmt"
)

type LessThanExpr struct {
	left  Expression
	right Expression
}

func LessThan(left Expression, right Expression) *LessThanExpr {
	l := new(LessThanExpr)
	l.left = left
	l.right = right
	return l
}

func (l LessThanExpr) String() string {
	return fmt.Sprintf("(%s < %s)", l.left, l.right)
}

func (l LessThanExpr) IsReducable() bool {
	return true
}

func (l LessThanExpr) Reduce(env *Environment) Expression {
	if l.left.IsReducable() {
		return LessThan(l.left.Reduce(env), l.right)
	} else if l.right.IsReducable() {
		return LessThan(l.left, l.right.Reduce(env))
	} else {
		left, ok := l.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := l.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Boolean(left < right)
		return res
	}
}

func (l LessThanExpr) Value() any {
	panic("Attempt to get value from LessThan")
}

func (l LessThanExpr) Evaluate(env *Environment) Expression {
	return Boolean(l.left.Evaluate(env).Value().(int) < l.right.Evaluate(env).Value().(int))
}
