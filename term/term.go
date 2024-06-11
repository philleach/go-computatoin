package term

type Term interface {
	Value() int
}
type Number struct {
	value int
}
type Multiply struct {
	left  Term
	right Term
}
type Add struct {
	left  Term
	right Term
}

func (n Number) Value() int {
	return n.value
}

func (m Multiply) Value() int {
	return m.left.Value() * m.right.Value()
}

func (a Add) Value() int {
	return a.left.Value() + a.right.Value()
}
