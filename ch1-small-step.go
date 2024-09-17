package main

import (
	"fmt"

	m "github.com/philleach/gocomputation/small-step/machine"
	s "github.com/philleach/gocomputation/small-step/syntax"
)

func main() {

	env := s.NewEnvironment()

	fmt.Println("*\n* Simple calc\n*")
	t2 := s.Multiply(s.Number(1), s.Number(3))
	m.Evaluator(t2, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Compound calc\n*")
	t3 := s.Add(s.Multiply(s.Number(1), s.Number(2)), s.Multiply(s.Number(3), s.Number(4)))
	m.Evaluator(t3, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc\n*")
	b := s.And(s.And(s.Boolean(true), s.Boolean(false)), s.Boolean(true))
	m.Evaluator(b, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Boolean calc with Not\n*")
	b2 := s.Not(s.And(s.Or(s.Boolean(true), s.Boolean(false)), s.Boolean(true)))
	m.Evaluator(b2, env).Run()
	fmt.Println("")

	fmt.Println("*\n* Equals\n*")
	eq := s.Equals(
		s.Add(s.Number(4), s.Multiply(s.Number(2), s.Number(4))),
		s.Subtract(s.Number(24), s.Divide(s.Number(24), s.Number(2))))
	m.Evaluator(eq, env).Run()
	fmt.Println("")

	fmt.Println("*\n* With comparitors\n*")
	x := s.Not(s.And(
		s.LessThan(s.Number(4), s.Number(5)),
		s.Not(s.GreaterThan(s.Number(2), s.Number(24)))))
	m.Evaluator(x, env).Run()
	fmt.Println("")

	fmt.Println("*\n* With variables\n*")

	env.Add("x", s.Number(6))
	env.Add("y", s.Number(4))
	z := s.Add(s.Variable("x"), s.Variable("y"))
	m.Evaluator(z, env).Run()
	fmt.Println("")

	aa := s.DoNothing()
	m.Machine(aa, env).Run()
	fmt.Println("")

	bb := s.Assign("x", s.Add(s.Number(4), s.Number(5)))
	m.Machine(bb, env).Run()
	fmt.Println("")

	e := s.NewEnvironment()
	e.Add("x", s.Number(0))
	st := s.Assign("x", s.Add(s.Variable("x"), s.Number(1)))
	m.Machine(st, e).Run()
	fmt.Println("")

	i := s.If(s.Equals(s.Variable("x"), s.Number(1)), s.Assign("y", s.Number(1)), s.Assign("y", s.Number(2)))
	m.Machine(i, e).Run()
	fmt.Println("")

	j := s.Sequence(s.Sequence(s.Assign("x", s.Add(s.Variable("x"), s.Number(1))),
		s.Assign("y", s.Add(s.Variable("x"), s.Number(2)))), s.Assign("z", s.Add(s.Variable("y"), s.Number(3))))
	m.Machine(j, e).Run()
	fmt.Println("")

	e.Add("x", s.Number(1))
	k := s.While(s.LessThan(s.Variable("x"), s.Number(5)), s.Assign("x", s.Multiply(s.Variable("x"), s.Number(3))))
	m.Machine(k, e).Run()
	fmt.Println("")
}
