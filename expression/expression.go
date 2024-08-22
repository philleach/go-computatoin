package expression

type Expression[T int | bool] interface {
	IsReducable() bool
	Reduce() Expression[T]
	Value() T
}
