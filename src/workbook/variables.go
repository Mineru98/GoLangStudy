package main

import "fmt"

// Go에서 변수는 명시적으로 선언되며 컴파일러에 의해 사용된다.
// 예를 들어, 함수 호출에서 타입이 정확한지 검사하는데 사용된다.

func main() {
	// var는 하나 또는 여러개의 변수를 선언한다.
	var a string = "initial"
	fmt.Println(a)

	// 한번에 여러개의 변수를 선언할 수 있다.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go는 초기화된 변수의 타입을 추론한다.
	var d = true
	fmt.Println(d)

	// 초기화 없이 선언된 변수는 제로값을 갖게 된다.
	// 예를 들어 int의 초기값은 0이다.
	var e int
	fmt.Println(e)

	// := 문법은 변수를 선언하는 동시에 초기화하기 위한 단축 표현식이다.
	// 이를 다른 방법으로 표현하자
	// var f string = "short"
	f := "short"
	fmt.Println(f)
}
