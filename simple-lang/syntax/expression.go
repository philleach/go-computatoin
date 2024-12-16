package syntax

import (
	"fmt"
)

// Interface
type Expression interface {
	IsReducable() bool
	Reduce(env *Environment) Expression
	Value() any
	Evaluate(env *Environment) Expression
}

//
// Primatives
//

// Boolean
type BooleanValue struct {
	value bool
}

func Boolean(v bool) *BooleanValue {
	b := new(BooleanValue)
	b.value = v
	return b
}

func (b BooleanValue) String() string {
	return fmt.Sprintf("%t", b.value)
}

func (b BooleanValue) Value() any {
	return b.value
}

// Number

type NumberValue struct {
	value int
}

func Number(i int) *NumberValue {
	n := new(NumberValue)
	n.value = i
	return n
}

func (n NumberValue) String() string {
	return fmt.Sprintf("%d", n.value)
}

func (n NumberValue) Value() any {
	return n.value
}

//
// Compounds
//

// Add
type AddExpr struct {
	left  Expression
	right Expression
}

func Add(left Expression, right Expression) *AddExpr {
	a := new(AddExpr)
	a.left = left
	a.right = right
	return a
}

func (a AddExpr) String() string {
	return fmt.Sprintf("(%s + %s)", a.left, a.right)
}

func (a AddExpr) Value() any {
	panic("Attempt to get value from Add")
}

// And
type AndExpr struct {
	left  Expression
	right Expression
}

func And(left Expression, right Expression) *AndExpr {
	a := new(AndExpr)
	a.left = left
	a.right = right
	return a
}

func (a AndExpr) String() string {
	return fmt.Sprintf("(%s && %s)", a.left, a.right)
}

func (a AndExpr) Value() any {
	panic("Attempt to get value from And")
}
