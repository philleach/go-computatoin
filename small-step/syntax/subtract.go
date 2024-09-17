package syntax

import (
	"fmt"
)

type SubtractExpr struct {
	left  Expression
	right Expression
}

func Subtract(left Expression, right Expression) *SubtractExpr {
	s := new(SubtractExpr)
	s.left = left
	s.right = right
	return s
}

func (s SubtractExpr) String() string {
	return fmt.Sprintf("(%s - %s)", s.left, s.right)
}

func (s SubtractExpr) IsReducable() bool {
	return true
}

func (s SubtractExpr) Reduce(env *Environment) Expression {
	if s.left.IsReducable() {
		return Subtract(s.left.Reduce(env), s.right)
	} else if s.right.IsReducable() {
		return Subtract(s.left, s.right.Reduce(env))
	} else {
		left, ok := s.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := s.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Number(left - right)
		return res
	}
}

func (s SubtractExpr) Value() any {
	panic("Attempt to get value from Add")
}
