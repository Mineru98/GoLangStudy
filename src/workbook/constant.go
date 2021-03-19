package main

import (
	"fmt"
	"math"
)

// const 키워드를 사용해 상수를 선안한다.
const c string = "constant"

func main() {
	fmt.Println(c)

	// const문은 var문과 동일하게 사용할 수 있다.
	const n = 50000000

	// 상수 표현식은 임의의 정밀도로 산술 연산을 수행한다.
	const d = 3e20 / n
	fmt.Println(d)

	// 숫자 상수는 명시적 캐스팅 등으로 타입이 주어지기 전까진 타입을 가지지 않는다.
	fmt.Println(int64(d))

	// 숫자는 변수 할당이나 함수 호출과 같은 context에서 사용하여 타입을 부여할 수 있다.
	// 예를 들면, math.Sin은 float64를 기대한다.
	fmt.Println(math.Sin(n))
}
