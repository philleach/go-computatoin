package syntax

import (
	"fmt"
)

type MultiplyExpr struct {
	left  Expression
	right Expression
}

func Multiply(left Expression, right Expression) *MultiplyExpr {
	m := new(MultiplyExpr)
	m.left = left
	m.right = right
	return m
}

func (m MultiplyExpr) String() string {
	return fmt.Sprintf("(%s * %s)", m.left, m.right)
}

func (m MultiplyExpr) IsReducable() bool {
	return true
}

func (m MultiplyExpr) Reduce(env *Environment) Expression {
	if m.left.IsReducable() {
		return Multiply(m.left.Reduce(env), m.right)
	} else if m.right.IsReducable() {
		return Multiply(m.left, m.right.Reduce(env))
	} else {
		left, ok := m.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := m.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		return Number(left * right)
	}
}

func (m MultiplyExpr) Value() any {
	panic("Attempt to get value from Multiply")
}

func (m MultiplyExpr) Evaluate(env *Environment) Expression {
	return Number(m.left.Evaluate(env).Value().(int) * m.right.Evaluate(env).Value().(int))
}
