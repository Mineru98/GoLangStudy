package main

import "fmt"

/*
	함수에서 결과값과 에러를 동시에 변환하는 경우에 사용한다.
*/

// 함수 시그니처에 있는 (int, int)는 이 함수가 2개의 int를 반환한다는걸 의미한다.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 다음은 함수 호출로부터 반환되는 두 반환값을
	// 다중 할당으로 받는다.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 반환값들 중 일부만 사용하고 싶을땐, 공백 식별자를 사용한다.
	_, c := vals()
	fmt.Println(c)
}
