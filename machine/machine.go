package machine

import (
	"fmt"

	"github.com/philleach/gocomputation/expression"
)

type Machine[T int | bool] struct {
	exp expression.Expression[T]
}

func NewMachine[T int | bool](exp expression.Expression[T]) *Machine[T] {
	m := new(Machine[T])
	m.exp = exp
	return m
}

func (m Machine[T]) Run() T {
	e := m.exp
	fmt.Printf("Evaluating : %s\n", e)

	for e.IsReducable() {
		e = e.Reduce()
		fmt.Printf("...is reduced to : %s\n", e)
	}
	return e.Value()
}
