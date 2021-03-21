package main

import "fmt"

/*
	closure를 만들 수 있는 익명 함수를 지원한다.
	익명 함수는 이름이 없는 한줄로 함수를 정의하고 싶을 때 유용하다.
*/

/*
	intSeq 함수는 내부에 익명으로 정의한 또 다른 함수를
	반환한다. 반환된 함수는 클로저를 만들기 위해 i변수를 가둬둔다.
*/
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	/*
		intSeq를 호출하여, 결과값을 nextInt에 할당한다.
		이 함수값은 nextInt를 호출할 때맏 업데이트 되는 i 값을 캡처한다.
	*/
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 특정 함수에서 상태값이 유일한진 확인하기 위해, 하나를 새롭게 생성하고 테스트를 해보자
	newInts := intSeq()
	fmt.Println(newInts())
}
