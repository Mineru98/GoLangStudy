package main

import "fmt"

// 두 int를 받아 그 합을 int 타입으로 반환하는 함수
func plus(a int, b int) int {
	return a + b
}

// Go에선 명시적 변환을 해야한다.
// 즉 마지막 식의 값을 자동으로 반환하지 않는다.

/*
	만약 같은 타입을 갖는 파라미터들이 연속적으로 주어지면,
	마지막 파라미터만 타입 입력을 선언하고 나머지 파라미터들에 대해서는 타입명을 생략 할 수 있다.
*/
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)

	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
