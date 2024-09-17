package syntax

import (
	"fmt"
)

type DivideExpr struct {
	left  Expression
	right Expression
}

func Divide(left Expression, right Expression) *DivideExpr {
	d := new(DivideExpr)
	d.left = left
	d.right = right
	return d
}

func (d DivideExpr) String() string {
	return fmt.Sprintf("(%s / %s)", d.left, d.right)
}

func (d DivideExpr) IsReducable() bool {
	return true
}

func (d DivideExpr) Reduce(env *Environment) Expression {
	if d.left.IsReducable() {
		return Divide(d.left.Reduce(env), d.right)
	} else if d.right.IsReducable() {
		return Divide(d.left, d.right.Reduce(env))
	} else {
		left, ok := d.left.Value().(int)
		if !ok {
			panic("Left value is not an int")
		}
		right, ok := d.right.Value().(int)
		if !ok {
			panic("Right value is not an int")
		}
		res := Number(left / right)
		return res
	}
}

func (d DivideExpr) Value() any {
	panic("Attempt to get value from Divide")
}
