package expression

import (
	"fmt"
)

type GreaterThan struct {
	left  Expression[int]
	right Expression[int]
}

func NewGreaterThan(left Expression[int], right Expression[int]) *GreaterThan {
	g := new(GreaterThan)
	g.left = left
	g.right = right
	return g
}

func (g GreaterThan) String() string {
	return fmt.Sprintf("(%s > %s)", g.left, g.right)
}

func (g GreaterThan) IsReducable() bool {
	return true
}

func (g GreaterThan) Reduce() Expression[bool] {
	if g.left.IsReducable() {
		return NewGreaterThan(g.left.Reduce(), g.right)
	} else if g.right.IsReducable() {
		return NewGreaterThan(g.left, g.right.Reduce())
	} else {
		res := NewBoolean(g.left.Value() > g.right.Value())
		return res
	}
}

func (g GreaterThan) Value() bool {
	panic("Attempt to get value from GreaterThan")
}
