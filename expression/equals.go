package expression

import (
	"fmt"
)

type Equals struct {
	left  Expression[int]
	right Expression[int]
}

func NewEquals(left Expression[int], right Expression[int]) *Equals {
	e := new(Equals)
	e.left = left
	e.right = right
	return e
}

func (a Equals) String() string {
	return fmt.Sprintf("%s = %s", a.left, a.right)
}

func (a Equals) IsReducable() bool {
	return true
}

func (a Equals) Reduce() Expression[bool] {
	if a.left.IsReducable() {
		return NewEquals(a.left.Reduce(), a.right)
	} else if a.right.IsReducable() {
		return NewEquals(a.left, a.right.Reduce())
	} else {
		res := NewBoolean(a.left.Value() == a.right.Value())
		return res
	}
}

func (a Equals) Value() bool {
	panic("Attempt to get value from Equals")
}
