package main

import (
	"fmt"
)

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

type calculator func(int, int) int

func t_calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	// 익명함수
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1,2,3,4,5)
	fmt.Println(result)
	fmt.Println()
	
	// 일급함수
	add := func(i int, j int) int {
		return i + j
	}
	
	r1 := calc(add, 10, 20)
	fmt.Println(r1)
	r2 := calc(func(x int, y int) int {return x - y}, 10, 20)
	fmt.Println(r2)
	fmt.Println()
	
	// type문을 사용한 함수 원형 정의
	r3 := t_calc(add, 30, 40)
	fmt.Println(r3)
	fmt.Println()
}