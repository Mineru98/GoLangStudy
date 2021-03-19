package main

import (
	"fmt"
)

func main() {
	// 변수
	var a int
	var f float32 = 3.14
	fmt.Println(a, f)
	a = 10
	f = 2.
	fmt.Println(a, f)
	var i, j, k int
	i, j, k = 1, 2, 3
	fmt.Println(i, j, k)
	// 상수
	const b int = 20
	const s string = "Hi"
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	fmt.Println(Visa, Master, Amex)
	// 키워드
	/*
		breack		default			func		interface		select
		case			defer				go			map					struct
		chan			else				goto		package			switch
		const			fallthrough	if			range				type
		continue	for					import	return			var
	*/
}
