package machine

import (
	"fmt"

	s "github.com/philleach/gocomputation/simple-lang/syntax"
)

type MachineDefn struct {
	statmt s.Statement
	env    *s.Environment
}

func Machine(statmt s.Statement, env *s.Environment) *MachineDefn {
	m := new(MachineDefn)
	m.statmt = statmt
	m.env = env
	return m
}

func (m MachineDefn) Run() {
	fmt.Printf("Evaluating       : %-40s %s\n", m.statmt, m.env)

	for m.statmt.IsReducable() {
		m.statmt, m.env = m.statmt.Reduce(m.env)
		fmt.Printf("...is reduced to : %-40s %s\n", m.statmt, m.env)
	}
	return
}
