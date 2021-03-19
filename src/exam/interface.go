package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// Rect 구조체에 대한 Shape 인터페이스 구현
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

// Circle 구조체에 대한 Shape 인터페이스 구현
func (c Circle) area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64 { return math.Pi * c.radius }

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		fmt.Println(a)
	}
}

func PrintIt(v interface{}) {
	fmt.Println(v)
}

func main() {
	// 인터페이스
	r := Rect{10, 20}
	c := Circle{10}
	showArea(r, c)
	fmt.Println()

	// 인터페이스 타입
	/*
		C#, Java에서 object 타입, c/c++에서 void* 타입
	*/
	var x interface{}
	x = 1
	PrintIt(x)
	x = "Test"
	PrintIt(x)
	fmt.Println()

	// Type AssertIon
	var a interface{} = 1
	i := a
	j := a.(int)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println()
}
