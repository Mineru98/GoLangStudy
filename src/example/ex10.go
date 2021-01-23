package main

import (
	"fmt"
)


func main() {
	// 슬라이스
	var arr []int
	arr = []int{1,2,3}
	arr[1] = 10
	fmt.Println(arr)
	
	arr2 := make([]int, 5, 10)
	fmt.Println(len(arr2), cap(arr2), arr2)// len은 길이, cap는 메모리 용량
	fmt.Println()
	
	var arr3 []int
	
	if arr3 == nil {
		fmt.Println("Nul Slice")
	}
	fmt.Println(len(arr3), cap(arr3))
	fmt.Println()
	
	// 부분 슬라이스
	arr4 := []int{0,1,2,3,4,5}
	arr4 = arr4[2:5]
	fmt.Println(arr4)
	fmt.Println()
	
	// 슬라이스 추가, 병합, 복사
	arrCopy := []int{0,1}
	arrCopy = append(arrCopy, 2)
	arrCopy = append(arrCopy, 3, 4, 5)
	fmt.Println(arrCopy)
	fmt.Println()
	
	sliceA := make([]int, 0, 3)
	
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)
	fmt.Println()
	
	sliceB := []int{1,2,3}
	sliceC := []int{4,5,6}
	
	sliceB = append(sliceB, sliceC...)
	fmt.Println(sliceB)
	fmt.Println()
	
	source := []int{7,8,9}
	target := make([]int, len(source), cap(source) * 2)
	copy(target, source)
	fmt.Println(len(target), cap(target), target)
	fmt.Println()
}