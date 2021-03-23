package main

import (
	"fmt"
	"os"
)

// Go는 Printf 전통에서 문자열 포맷팅을 훌룡하게 제공한다.

type point struct {
	x, y int
}

func main() {
	// Go는 일반적인 값들을 포맷팅 하기 위해 설계된 몇 가지 출력 "verbs"를 제공한다.
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// 값이 구조체인 경우, %+v 변형은 구조체의 필드명까지 포함한다.
	fmt.Printf("%+v\n", p)
	// %#v 변형은 값의 Go 구문 표현 즉, 해당 값을 생성하는 소스 코드 스니펫을 출력한다.
	fmt.Printf("%#v\n", p)

	// 값의 타입을 출력하기 위해선, %T를 사용한다.
	fmt.Printf("%T\n", p)
	// 불리언 포맷팅은 직관적이다.
	fmt.Printf("%t\n", true)
	// 정수 포맷팅에서는 여러가지 옵션이 있다.
	// 표준적인 10진수 포맷팅은 %d를 사용한다.
	fmt.Printf("%d\n", 123)

	// 바이너리 표현 출력
	fmt.Printf("%b\n", 14)
	// 주어진 정수에 해당하는 문자를 출력
	fmt.Printf("%c\n", 33)
	// %x는 16진수 인코딩을 제공
	fmt.Printf("%x\n", 456)
	// 실수에 대해서도 몇가지 포맷팅 옵션이 있다.
	// 기본적인 십진수 포맷팅에 대해서는 %f를 사용한다.
	fmt.Printf("%f\n", 78.9)

	// %e와 %E(약간 다른 버전)는 실수를 과학적 표기법으로 포맷팅한다.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 기본적인 문자열 출력에는 %s를 사용한다.
	fmt.Printf("%s\n", "\"string\"")
	// Go 소스에서와 같이 문자열을 쌍따옴표로 묶으려면, %q를 사용한다.
	fmt.Printf("%q\n", "\"string\"")

	// 앞에서 본 정수와 마찬가지로, x는 입력값이 바이트당 두 개의
	// 출력문자와 함께 문자열을 16진수로 렌더링한다.
	fmt.Printf("%x\n", "hex this")
	// 포인터의 표현(메모리 표현)을 출력하기 위해선, %p를 사용한다.
	fmt.Printf("%p\n", &p)
	/*
		숫자를 포맷팅할때 결과 형태의 너비와 정확도를 조정하고 싶을때가 종종 있다.
		정수의 너비를 지정하려면, 숫자를 % 다음에 쓴다.
		기본값으로 결과는 우측 정렬이며 공백으로 채워진다.
	*/
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// 출력된 실수에 너비를 지정할 수도 있지만, 일반적으로는
	// windth.precision 구문을 사용해 정밀도를 제한하고자 한다.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// 좌측 정렬을 위해선, - 플래그를 사용한다.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// 문자열을 포맷팅할 때, 특히 테이블같은 형태로 정렬하기 위해
	// 너비를 조정하고 싶을때도 있다. 다음은 기본적인 우측 정렬 너비의 예시이다.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	// 이제까지 우린 os.Stdout에 포맷팅된 문자열을 출력하는 Printf를 사용해봤다.
	// Sprintf는 문자열을 출력하지 않고 이를 포맷팅 한 후 반환한다.
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	// Fprintf를 사용하여 os.Stdout 이외의 다른 io.Writers를 포맷팅하고 출력할 수 있다.
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
