package expression

import (
	"fmt"
)

type BooleanValue struct {
	value bool
}

func Boolean(v bool) *BooleanValue {
	b := new(BooleanValue)
	b.value = v
	return b
}

func (b BooleanValue) String() string {
	return fmt.Sprintf("%t", b.value)
}

func (b BooleanValue) IsReducable() bool {
	return false
}

func (b BooleanValue) Reduce(env *Environment) Expression {
	panic("Attempt to reduce a Boolean")
}

func (b BooleanValue) Value() any {
	return b.value
}
