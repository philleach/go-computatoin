package main

import (
	"fmt"

	e "github.com/philleach/gocomputation/expression"
	m "github.com/philleach/gocomputation/machine"
)

func main() {

	env := e.NewEnvironment()

	fmt.Println("*\n* Simple calc\n*")
	t2 := e.Multiply(e.Number(1), e.Number(3))
	m.Machine(t2, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Compound calc\n*")
	t3 := e.Add(e.Multiply(e.Number(1), e.Number(2)), e.Multiply(e.Number(3), e.Number(4)))
	m.Machine(t3, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc\n*")
	b := e.And(e.And(e.Boolean(true), e.Boolean(false)), e.Boolean(true))
	m.Machine(b, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc with Not\n*")
	b2 := e.Not(e.And(e.Or(e.Boolean(true), e.Boolean(false)), e.Boolean(true)))
	m.Machine(b2, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Equals\n*")
	eq := e.Equals(
		e.Add(e.Number(4), e.Multiply(e.Number(2), e.Number(4))),
		e.Subtract(e.Number(24), e.Divide(e.Number(24), e.Number(2))))
	m.Machine(eq, env).Run()
	fmt.Println("")

	fmt.Println("*\n* With comparitors\n*")
	x := e.Not(e.And(
		e.LessThan(e.Number(4), e.Number(5)),
		e.Not(e.GreaterThan(e.Number(2), e.Number(24)))))
	m.Machine(x, env).Run()
	fmt.Println("")

	fmt.Println("*\n* With variables\n*")

	env.Add("x", e.Number(6))
	env.Add("y", e.Number(4))
	z := e.Add(e.Variable("x"), e.Variable("y"))
	m.Machine(z, env).Run()
	fmt.Println("")
}
