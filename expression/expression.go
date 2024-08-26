package expression

type Expression[T int | bool] interface {
	IsReducable() bool
	Reduce(env *Environment) Expression[T]
	Value() T
}

type Environment map[string]any
