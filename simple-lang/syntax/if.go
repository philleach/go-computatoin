package syntax

import (
	"fmt"
)

type IfStmt struct {
	condition   Expression
	consequence Statement
	alternative Statement
}

func If(cond Expression, cons Statement, alt Statement) *IfStmt {
	i := new(IfStmt)
	i.condition = cond
	i.consequence = cons
	i.alternative = alt
	return i
}

func (i IfStmt) String() string {
	return fmt.Sprintf("if (%s) {%s} else {%s}", i.condition, i.consequence, i.alternative)
}

func (a IfStmt) IsReducable() bool {
	return true
}

func (i IfStmt) Reduce(env *Environment) (Statement, *Environment) {
	if i.condition.IsReducable() {
		return If(i.condition.Reduce(env), i.consequence, i.alternative), env
	} else {
		c, ok := i.condition.Value().(bool)
		if !ok {
			panic("Condition is not an bool")
		}
		if c {
			return i.consequence, env
		} else {
			return i.alternative, env
		}
	}
}

func (i IfStmt) Evaluate(env *Environment) *Environment {
	if i.condition.Evaluate(env).Value().(bool) {
		return i.consequence.Evaluate(env)
	} else {
		return i.alternative.Evaluate(env)
	}
}
