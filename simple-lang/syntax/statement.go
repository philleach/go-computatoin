package syntax

import "fmt"

type Statement interface {
	IsReducable() bool
	Reduce(env *Environment) (Statement, *Environment)
	Evaluate(env *Environment) *Environment
}

// Assign

type AssignStmt struct {
	name string
	exp  Expression
}

func Assign(name string, exp Expression) *AssignStmt {
	as := new(AssignStmt)
	as.name = name
	as.exp = exp
	return as
}

func (a AssignStmt) String() string {
	return fmt.Sprintf("%s := %s", a.name, a.exp)
}
