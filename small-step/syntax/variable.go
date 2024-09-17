package syntax

import (
	"fmt"
)

type VariableExpr struct {
	name string
}

func Variable(name string) *VariableExpr {
	v := new(VariableExpr)
	v.name = name
	return v
}

func (v VariableExpr) String() string {
	return fmt.Sprintf("%s", v.name)
}

func (v VariableExpr) IsReducable() bool {
	return true
}

func (v VariableExpr) Reduce(env *Environment) Expression {
	return env.lookup(v.name)
}

func (v VariableExpr) Value() any {
	panic("Attempt to get value from Equals")
}
