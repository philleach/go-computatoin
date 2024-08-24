package main

import (
	"fmt"

	e "github.com/philleach/gocomputation/expression"
	m "github.com/philleach/gocomputation/machine"
)

func main() {

	fmt.Println("*\n* Simple calc\n*")
	t2 := e.NewMultiply(e.NewNumber(1), e.NewNumber(3))
	m.NewMachine(t2).Run()
	fmt.Println("")

	fmt.Println("*\n* Compound calc\n*")
	t3 := e.NewAdd(e.NewMultiply(e.NewNumber(1), e.NewNumber(2)), e.NewMultiply(e.NewNumber(3), e.NewNumber(4)))
	m.NewMachine(t3).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc\n*")
	b := e.NewAnd(e.NewAnd(e.NewBoolean(true), e.NewBoolean(false)), e.NewBoolean(true))
	m.NewMachine(b).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc with Not\n*")
	b2 := e.NewNot(e.NewAnd(e.NewOr(e.NewBoolean(true), e.NewBoolean(false)), e.NewBoolean(true)))
	m.NewMachine(b2).Run()
	fmt.Println("")

	fmt.Println("*\n* Equals\n*")
	eq := e.NewEquals(
		e.NewAdd(e.NewNumber(4), e.NewMultiply(e.NewNumber(2), e.NewNumber(4))),
		e.NewSubtract(e.NewNumber(24), e.NewDivide(e.NewNumber(24), e.NewNumber(2))))
	m.NewMachine(eq).Run()
	fmt.Println("")

	fmt.Println("*\n* With comparitors\n*")
	x := e.NewNot(e.NewAnd(
		e.NewLessThan(e.NewNumber(4), e.NewNumber(5)),
		e.NewNot(e.NewGreaterThan(e.NewNumber(2), e.NewNumber(24)))))
	m.NewMachine(x).Run()
	fmt.Println("")
}
