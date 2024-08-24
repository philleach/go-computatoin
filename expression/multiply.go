package expression

import (
	"fmt"
)

type Multiply struct {
	left  Expression[int]
	right Expression[int]
}

func NewMultiply(left Expression[int], right Expression[int]) *Multiply {
	m := new(Multiply)
	m.left = left
	m.right = right
	return m
}

func (m Multiply) String() string {
	return fmt.Sprintf("(%s * %s)", m.left, m.right)
}

func (m Multiply) IsReducable() bool {
	return true
}

func (m Multiply) Reduce() Expression[int] {
	if m.left.IsReducable() {
		return NewMultiply(m.left.Reduce(), m.right)
	} else if m.right.IsReducable() {
		return NewMultiply(m.left, m.right.Reduce())
	} else {
		return NewNumber(m.left.Value() * m.right.Value())
	}
}

func (m Multiply) Value() int {
	panic("Attempt to get value from Multiply")
}
