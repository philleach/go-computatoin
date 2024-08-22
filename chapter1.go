package main

import (
	"fmt"

	"github.com/philleach/gocomputation/expression"
	"github.com/philleach/gocomputation/machine"
)

func main() {
	t := expression.NewNumber(5)
	t2 := expression.NewMultiply(expression.NewNumber(1), expression.NewNumber(3))
	t3 := expression.NewAdd(expression.NewMultiply(expression.NewNumber(1), expression.NewNumber(2)), expression.NewMultiply(expression.NewNumber(3), expression.NewNumber(4)))

	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println("")

	machine.NewMachine(t3).Run()

}
