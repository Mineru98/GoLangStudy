package main

import (
	"fmt"
)

func main() {
	// 배열
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)
	fmt.Println(a[2])
	fmt.Println()

	// 배열 초기화
	var arr1 = [3]int{11, 22, 33}
	var arr2 = [...]int{44, 55, 66}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println()

	// 다차원 배열
	var multiArr [3][4][5]int
	multiArr[0][0][0] = 10
	fmt.Println(multiArr)
	fmt.Println()

	// 다차원 배열 초기화
	var multiArrInit = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(multiArrInit)
}
