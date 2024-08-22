package expression

import (
	"fmt"
)

type Divide struct {
	left  Expression[int]
	right Expression[int]
}

func NewDivide(left Expression[int], right Expression[int]) *Divide {
	d := new(Divide)
	d.left = left
	d.right = right
	return d
}

func (d Divide) String() string {
	return fmt.Sprintf("%s + %s", d.left, d.right)
}

func (d Divide) IsReducable() bool {
	return true
}

func (d Divide) Reduce() Expression[int] {
	if d.left.IsReducable() {
		return NewDivide(d.left.Reduce(), d.right)
	} else if d.right.IsReducable() {
		return NewDivide(d.left, d.right.Reduce())
	} else {
		res := NewNumber(d.left.Value() / d.right.Value())
		return res
	}
}

func (d Divide) Value() int {
	panic("Attempt to get value from Divide")
}
