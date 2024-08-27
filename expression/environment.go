package expression

import (
	"fmt"
)

type Environment struct {
	env map[string]Expression
}

func NewEnvironment() *Environment {
	e := new(Environment)
	e.env = make(map[string]Expression)
	return e
}

func (e Environment) Add(name string, exp Expression) {
	e.env[name] = exp
}

func (e Environment) String() string {
	return fmt.Sprintf("Environment")
}

func (e Environment) lookup(name string) Expression {
	return e.env[name]
}
