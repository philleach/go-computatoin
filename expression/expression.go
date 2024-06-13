package expression

type Expression interface {
	IsReducable() bool
	Reduce() Expression
	Value() int
}
