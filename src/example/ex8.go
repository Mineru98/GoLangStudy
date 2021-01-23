package main

import (
	"fmt"
)

func nextVal() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 클로저
	next := nextVal()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	
	nextOther := nextVal()
	fmt.Println(nextOther())
	fmt.Println(nextOther())
	fmt.Println()
}