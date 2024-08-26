package machine

import (
	"fmt"

	e "github.com/philleach/gocomputation/expression"
)

type Machine[T int | bool] struct {
	exp e.Expression[T]
}

func NewMachine[T int | bool](exp e.Expression[T]) *Machine[T] {
	m := new(Machine[T])
	m.exp = exp
	return m
}

func (m Machine[T]) Run() T {
	exp := m.exp
	fmt.Printf("Evaluating : %s\n", exp)

	env := new(e.Environment)
	for exp.IsReducable() {
		exp = exp.Reduce(env)
		fmt.Printf("...is reduced to : %s\n", exp)
	}
	return exp.Value()
}
