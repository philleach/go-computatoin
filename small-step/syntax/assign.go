package syntax

import (
	"fmt"
)

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

func (a AssignStmt) IsReducable() bool {
	return true
}

func (a AssignStmt) Reduce(env *Environment) (Statement, *Environment) {
	if a.exp.IsReducable() {
		return Assign(a.name, a.exp.Reduce(env)), env
	} else {
		return DoNothing(), env.Merge(a.name, a.exp)
	}
}
