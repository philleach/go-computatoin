package machine

import (
	"fmt"

	s "github.com/philleach/gocomputation/small-step/syntax"
)

type EvaluatorDefn struct {
	exp s.Expression
	env *s.Environment
}

func Evaluator(exp s.Expression, env *s.Environment) *EvaluatorDefn {
	m := new(EvaluatorDefn)
	m.exp = exp
	m.env = env
	return m
}

func (m EvaluatorDefn) Run() {
	fmt.Printf("Evaluating       : %-40s %s\n", m.exp, m.env)

	for m.exp.IsReducable() {
		m.exp = m.exp.Reduce(m.env)
		fmt.Printf("...is reduced to : %-40s %s\n", m.exp, m.env)
	}
	return
}
