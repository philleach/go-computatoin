package expression

import (
	"fmt"
)

type Number struct {
	value int
}

func NewNumber(i int) *Number {
	n := new(Number)
	n.value = i
	return n
}

func (n Number) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n Number) IsReducable() bool {
	return false
}

func (n Number) Reduce() Expression {
	panic("Attempt to reduce a Number")
}

func (n Number) Value() int {
	return n.value
}
