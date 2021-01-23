package main

import (
	"fmt"
)

func call(msg string) {
	fmt.Println(msg)
}

func passByRef(msg *string) {
	fmt.Println(*msg)
	*msg = "Luck!"
}

func varFunc(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func twiceSum(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func twiceSumOther(nums ...int) (count int, s int) {
	for _, n := range nums {
		s += n
	}
	count = len(nums)
	return
}

func main() {
	// 함수
	call("Hello, World!")
	fmt.Println()
	
	// Pass By Reference
	msg := "Good"
	passByRef(&msg)
	fmt.Println(msg)
	fmt.Println()
	
	// 가변인자함수
	varFunc("I", "Love", "You")
	
	// 함수 변환값
	total := sum(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(total)
	fmt.Println()
	
	count1, total1 := twiceSum(1,2,3,4,5)
	fmt.Println(count1, total1)
	fmt.Println()
	
	count2, total2 := twiceSumOther(1,2,3)
	fmt.Println(count2, total2)
	fmt.Println()
}