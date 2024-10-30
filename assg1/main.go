package main

import "fmt"

type Triangle struct {
	height float64
	base   float64
}

type Square struct {
	sideLength float64
}

func (t Triangle) area() float64 {
	return 0.5 * t.height * t.base
}
func (s Square) area() float64 {
	return s.sideLength * s.sideLength
}

type Shape interface {
	area() float64
}

func printArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	fmt.Println("Hello, World!")
	triangle := Triangle{height: 4.0, base: 6.0}
	square := Square{sideLength: 5.0}
	fmt.Println("Area of Triangle:", printArea(triangle), printArea(square))
}
