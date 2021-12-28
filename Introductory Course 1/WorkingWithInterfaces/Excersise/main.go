package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type triangle struct {
	length float64
	base   float64
	height float64
}

type sqaure struct {
	length float64
}

func main() {
	triangleInstance := triangle{length: 2.0, base: 2.0, height: 3.0}
	squareInstance := sqaure{length: 4.0}
	printArea(triangleInstance)
	printArea(squareInstance)

}

func (s sqaure) getArea() float64 {
	return math.Exp2(s.length)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
