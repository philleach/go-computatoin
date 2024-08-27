package expression

type Expression interface {
	IsReducable() bool
	Reduce(env *Environment) Expression
	Value() any
}
