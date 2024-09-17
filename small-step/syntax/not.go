package syntax

import (
	"fmt"
)

type NotExpr struct {
	value Expression
}

func Not(value Expression) *NotExpr {
	n := new(NotExpr)
	n.value = value
	return n
}

func (n NotExpr) String() string {
	return fmt.Sprintf("!(%s)", n.value)
}

func (n NotExpr) IsReducable() bool {
	return true
}

func (n NotExpr) Reduce(env *Environment) Expression {
	if n.value.IsReducable() {
		return Not(n.value.Reduce(env))
	} else {
		value, ok := n.value.Value().(bool)
		if !ok {
			panic("Left value is not an bool")
		}
		res := Boolean(!value)
		return res
	}
}

func (n NotExpr) Value() any {
	panic("Attempt to get value from Not")
}
