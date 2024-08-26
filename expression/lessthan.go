package expression

import (
	"fmt"
)

type LessThan struct {
	left  Expression[int]
	right Expression[int]
}

func NewLessThan(left Expression[int], right Expression[int]) *LessThan {
	l := new(LessThan)
	l.left = left
	l.right = right
	return l
}

func (l LessThan) String() string {
	return fmt.Sprintf("(%s < %s)", l.left, l.right)
}

func (l LessThan) IsReducable() bool {
	return true
}

func (l LessThan) Reduce(env *Environment) Expression[bool] {
	if l.left.IsReducable() {
		return NewLessThan(l.left.Reduce(env), l.right)
	} else if l.right.IsReducable() {
		return NewLessThan(l.left, l.right.Reduce(env))
	} else {
		res := NewBoolean(l.left.Value() < l.right.Value())
		return res
	}
}

func (l LessThan) Value() bool {
	panic("Attempt to get value from LessThan")
}
