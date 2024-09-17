package syntax

import (
	"fmt"
)

type DoNothingStmt struct {
}

func DoNothing() *DoNothingStmt {
	dn := new(DoNothingStmt)
	return dn
}

func (a DoNothingStmt) String() string {
	return fmt.Sprintf("do-nothing")
}

func (a DoNothingStmt) IsReducable() bool {
	return false
}

func (a DoNothingStmt) Reduce(env *Environment) (Statement, *Environment) {
	panic("Attemmpt to reduce DoNothing Statament")
}
