package main

import "fmt"

// Go에는 string, int, float, bool 등을 포함한 다양한 타입들이 존재한다.

func main() {

	// 문자열은 +로 합칠 수 있다.
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
