package expression

import (
	"fmt"
)

type NumberValue struct {
	value int
}

func Number(i int) *NumberValue {
	n := new(NumberValue)
	n.value = i
	return n
}

func (n NumberValue) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n NumberValue) IsReducable() bool {
	return false
}

func (n NumberValue) Reduce(env *Environment) Expression {
	panic("Attempt to reduce a Number")
}

func (n NumberValue) Value() any {
	return n.value
}
