package main

import (
	"fmt"
	in "github.com/yfedoruck/gopattern/pkg/interpreter"
)

func main() {

	calc := in.Calculator{}
	add := calc.Interpret("one + three")
	sub := calc.Interpret("three - two")
	mul := calc.Interpret("three * two")
	fmt.Println(add == 4)
	fmt.Println(sub == 1)
	fmt.Println(mul == 6)
}
