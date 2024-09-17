package syntax

import (
	"fmt"
	"maps"
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

func (e Environment) Merge(name string, exp Expression) *Environment {
	ne := new(Environment)

	ne.env = maps.Clone(e.env)
	ne.env[name] = exp
	return ne
}

func (e Environment) String() string {
	var s string = "{ "
	for key, val := range e.env {
		s = s + fmt.Sprintf(":%s=>«%s» ", key, val)
	}
	s = s + "}"
	return s
}

func (e Environment) lookup(name string) Expression {
	return e.env[name]
}
