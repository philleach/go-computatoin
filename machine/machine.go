package machine

import (
	"fmt"

	e "github.com/philleach/gocomputation/expression"
)

type MachineDefn struct {
	exp e.Expression
	env *e.Environment
}

func Machine(exp e.Expression, env *e.Environment) *MachineDefn {
	m := new(MachineDefn)
	m.exp = exp
	m.env = env
	return m
}

func (m MachineDefn) Run() any {
	exp := m.exp
	fmt.Printf("Evaluating : %s\n", exp)

	for exp.IsReducable() {
		exp = exp.Reduce(m.env)
		fmt.Printf("...is reduced to : %s\n", exp)
	}
	return exp.Value()
}
