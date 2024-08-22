package expression

import (
	"fmt"
)

type Boolean struct {
	value bool
}

func NewBoolean(v bool) *Boolean {
	b := new(Boolean)
	b.value = v
	return b
}

func (b Boolean) String() string {
	return fmt.Sprintf("%t", b.value)
}

func (b Boolean) IsReducable() bool {
	return false
}

func (b Boolean) Reduce() Expression[bool] {
	panic("Attempt to reduce a Boolean")
}

func (b Boolean) Value() bool {
	return b.value
}
