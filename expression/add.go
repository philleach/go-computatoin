package expression

import (
	"fmt"
)

type Add struct {
	left  Expression
	right Expression
}

func NewAdd(left Expression, right Expression) *Add {
	a := new(Add)
	a.left = left
	a.right = right
	return a
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.left, a.right)
}

func (a Add) IsReducable() bool {
	return true
}

func (a Add) Reduce() Expression {
	if a.left.IsReducable() {
		return NewAdd(a.left.Reduce(), a.right)
	} else if a.right.IsReducable() {
		return NewAdd(a.left, a.right.Reduce())
	} else {
		return NewNumber(a.left.Value() * a.right.Value())
	}
}

func (a Add) Value() int {
	panic("Attempt to get value from Add")
}
