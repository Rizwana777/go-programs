package main

import "fmt"

type Calculator interface {
	Addition()
	Subtraction()
}

type BasicCalculator struct {
	num1 int64
	num2 int64
}

// We can enforce interface so that we can make sure our structures uses interface.
func NewCalculator(arg int64, arg2 int64) Calculator {
	return &BasicCalculator{arg, arg2}
}

type ScientificCalculator struct {
	num1 int64
	num2 int64
}

func (b *BasicCalculator) Addition() {
	fmt.Println("Addition in basic calculator", b.num1+b.num2)
}

func (b *BasicCalculator) Subtraction() {
	fmt.Println("Subtraction in basic calculator", b.num1-b.num2)
}

func (s *ScientificCalculator) Addition() {
	fmt.Println("Addition in scientific calculator", s.num1+s.num2)
}

func (s *ScientificCalculator) Subtraction() {
	fmt.Println("Subtraction in scientific calculator", s.num1-s.num2)
}

func main() {
	b := BasicCalculator{50, 40}
	s := ScientificCalculator{50, 20}
	b.Addition()
	s.Addition()
	b.Subtraction()
}
