package main

import "fmt"

/*
	배열은 크기가 고정되어 있다.
	베열은 복사를 하면 별도 메모리를 생성한다.
*/
func main() {
	/*
		정확히 5개의 정수를 가진 arr라는 배열을 만든다.
		원소의 타입과 길이는 배열 타입의 두 부분이다.
		배열의 기본값은 0이며, int 배열의 경우 0 배열으 의미한다.
	*/
	var arr [5]int
	fmt.Println("emp:", arr)

	/*
		array[index] = value 문법을 사용해 특정 위치에 값을
		설정할 수 있으며, array[index]로 깂을 가져올 수 있다.
	*/
	arr[4] = 100
	fmt.Println("set:", arr)
	fmt.Println("get:", arr[4])

	// 내장 함수인 len은 배열의 길이를 반환한다.
	fmt.Println("len:", len(arr))

	// 한 줄에서 선언과 동시에 배열을 초기화 할때에는 다음과 같은 문법으로 한다.
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", arr2)

	// 배열 타입은 1차원이지만, 타입들을 구성하여
	// 다차원 구조를 만들 수 있다.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
