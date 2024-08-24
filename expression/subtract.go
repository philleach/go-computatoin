package expression

import (
	"fmt"
)

type Subtract struct {
	left  Expression[int]
	right Expression[int]
}

func NewSubtract(left Expression[int], right Expression[int]) *Subtract {
	s := new(Subtract)
	s.left = left
	s.right = right
	return s
}

func (s Subtract) String() string {
	return fmt.Sprintf("(%s - %s)", s.left, s.right)
}

func (s Subtract) IsReducable() bool {
	return true
}

func (s Subtract) Reduce() Expression[int] {
	if s.left.IsReducable() {
		return NewSubtract(s.left.Reduce(), s.right)
	} else if s.right.IsReducable() {
		return NewSubtract(s.left, s.right.Reduce())
	} else {
		res := NewNumber(s.left.Value() - s.right.Value())
		return res
	}
}

func (s Subtract) Value() int {
	panic("Attempt to get value from Add")
}
