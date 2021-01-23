package main

import (
	"fmt"
)

func main() {
	// 산술연산자
	var a, b, c int = 1, 2, 3
	c = (a + b) / 3
	c++;
	fmt.Println(c)
	fmt.Println()
	
	// 관계연산자
	fmt.Println(a == b)
	fmt.Println(a != c)
	fmt.Println(a >= b)
	fmt.Println()
	
	// 논리연산자
	fmt.Println(true && false)
	fmt.Println(true || !(false && true))
	fmt.Println()
	
	// Bitwise 연산자
	c = (a & b) << 5
	fmt.Println(c)
	fmt.Println()
	
	// 할당연산자
	a = 100
	a *= 10
	a >>= 2
	a |= 1
	fmt.Println(a)
	fmt.Println()
	
	// 포인터연산자
	var k int = 10
	var p = &k
	fmt.Println(*p, p)
	fmt.Println()
}