package syntax

import (
	"fmt"
)

type WhileStmt struct {
	cond Expression
	body Statement
}

func While(cond Expression, body Statement) *WhileStmt {
	w := new(WhileStmt)
	w.cond = cond
	w.body = body
	return w
}

func (w WhileStmt) String() string {
	return fmt.Sprintf("while (%s) { %s }", w.cond, w.body)
}

func (w WhileStmt) IsReducable() bool {
	return true
}

func (w WhileStmt) Reduce(env *Environment) (Statement, *Environment) {
	return If(w.cond,
		Sequence(w.body, While(w.cond, w.body)),
		DoNothing()), env
}

func (w WhileStmt) Evaluate(env *Environment) *Environment {
	if w.cond.Evaluate(env).Value().(bool) {
		env = w.body.Evaluate(env)
	} else {
		return env
	}
	return env
}
