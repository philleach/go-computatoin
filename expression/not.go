package expression

import (
	"fmt"
)

type Not struct {
	value Expression[bool]
}

func NewNot(value Expression[bool]) *Not {
	n := new(Not)
	n.value = value
	return n
}

func (n Not) String() string {
	return fmt.Sprintf("!(%s)", n.value)
}

func (n Not) IsReducable() bool {
	return true
}

func (n Not) Reduce(env *Environment) Expression[bool] {
	if n.value.IsReducable() {
		return NewNot(n.value.Reduce(env))
	} else {
		res := NewBoolean(!n.value.Value())
		return res
	}
}

func (n Not) Value() bool {
	panic("Attempt to get value from Not")
}
