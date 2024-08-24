package expression

import (
	"fmt"
)

type Or struct {
	left  Expression[bool]
	right Expression[bool]
}

func NewOr(left Expression[bool], right Expression[bool]) *Or {
	o := new(Or)
	o.left = left
	o.right = right
	return o
}

func (o Or) String() string {
	return fmt.Sprintf("(%s || %s)", o.left, o.right)
}

func (o Or) IsReducable() bool {
	return true
}

func (o Or) Reduce() Expression[bool] {
	if o.left.IsReducable() {
		return NewOr(o.left.Reduce(), o.right)
	} else if o.right.IsReducable() {
		return NewOr(o.left, o.right.Reduce())
	} else {
		res := NewBoolean(o.left.Value() || o.right.Value())
		return res
	}
}

func (o Or) Value() bool {
	panic("Attempt to get value from Or")
}
