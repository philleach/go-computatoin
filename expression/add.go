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
		fmt.Println("1")
		return NewAdd(a.left.Reduce(), a.right)
	} else if a.right.IsReducable() {
		fmt.Println("2")
		return NewAdd(a.left, a.right.Reduce())
	} else {
		fmt.Println("3")
		fmt.Printf("Left %v Right %v", a.left.Value(), a.right.Value())
		res := NewNumber(a.left.Value() + a.right.Value())
		fmt.Printf("Res %s ", res)
		return res
	}
}

func (a Add) Value() int {
	panic("Attempt to get value from Add")
}
