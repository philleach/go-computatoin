package syntax

type Expression interface {
	IsReducable() bool
	Reduce(env *Environment) Expression
	Value() any
}
type Statement interface {
	IsReducable() bool
	Reduce(env *Environment) (Statement, *Environment)
}
