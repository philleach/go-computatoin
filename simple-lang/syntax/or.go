package syntax

import (
	"fmt"
)

type OrExpr struct {
	left  Expression
	right Expression
}

func Or(left Expression, right Expression) *OrExpr {
	o := new(OrExpr)
	o.left = left
	o.right = right
	return o
}

func (o OrExpr) String() string {
	return fmt.Sprintf("(%s || %s)", o.left, o.right)
}

func (o OrExpr) IsReducable() bool {
	return true
}

func (o OrExpr) Reduce(env *Environment) Expression {
	if o.left.IsReducable() {
		return Or(o.left.Reduce(env), o.right)
	} else if o.right.IsReducable() {
		return Or(o.left, o.right.Reduce(env))
	} else {
		left, ok := o.left.Value().(bool)
		if !ok {
			panic("Left value is not an bool")
		}
		right, ok := o.right.Value().(bool)
		if !ok {
			panic("Right value is not an bool")
		}
		res := Boolean(left || right)
		return res
	}
}

func (o OrExpr) Value() any {
	panic("Attempt to get value from Or")
}

func (o OrExpr) Evaluate(env *Environment) Expression {
	return Boolean(o.left.Evaluate(env).Value().(bool) || o.right.Evaluate(env).Value().(bool))
}
