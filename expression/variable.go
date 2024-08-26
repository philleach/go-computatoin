package expression

import (
	"fmt"
)

type Variable struct {
	name string
}

func NewVariable(name string) *Variable {
	v := new(Variable)
	v.name = name
	return v
}

func (v Variable) String() string {
	return fmt.Sprintf("%s", v.name)
}

func (v Variable) IsReducable() bool {
	return true
}

func (v Variable) Reduce(env *Environment) any {
	return (*env)[v.name]
}

func (v Variable) Value() bool {
	panic("Attempt to get value from Equals")
}
