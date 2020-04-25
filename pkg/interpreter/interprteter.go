package interpreter

import "strings"

type Expression interface {
	Interpret() int
}
type Operation interface {
	Expression
	Left(Expression)
	Right(Expression)
}

type Calculator struct {
}

func (r Calculator) Interpret(s string) int {
	str := strings.Split(s, " ")

	leftExp := Number{Input: str[0]}
	operatorExp := Operator{Input: str[1]}
	rightExp := Number{Input: str[2]}

	operation := operatorExp.Operation()
	operation.Left(leftExp)
	operation.Right(rightExp)
	return operation.Interpret()
}

type Number struct {
	Input string
}

func (r Number) Interpret() int {
	switch r.Input {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	}

	return 0
}

type Operator struct {
	Input string
}

func (r Operator) Operation() Operation {
	switch r.Input {
	case "+":
		return &Add{}
	case "-":
		return &Sub{}
	case "*":
		return &Mul{}
	case "/":
		return &Div{}
	}

	return &Add{}
}

type Operand struct {
	left  Expression
	right Expression
}

func (r *Operand) Left(e Expression) {
	r.left = e
}
func (r *Operand) Right(e Expression) {
	r.right = e
}

//
type Add struct {
	Operand
}

func (r Add) Interpret() int {
	return r.left.Interpret() + r.right.Interpret()
}

//
type Sub struct {
	Operand
}

func (r Sub) Interpret() int {
	return r.left.Interpret() - r.right.Interpret()
}

//
type Div struct {
	Operand
}

func (r Div) Interpret() int {
	return r.left.Interpret() / r.right.Interpret()
}

//
type Mul struct {
	Operand
}

func (r Mul) Interpret() int {
	return r.left.Interpret() * r.right.Interpret()
}
