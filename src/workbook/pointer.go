package main

import "fmt"

/*
	두 함수 zeroval와 zeroptr를 가지고 포인터가 값과 다르게 어떻게 작동할까?

	zeroval은 int 타입 파라미터를 가지므로 인자는 값으로 전달된다.
	zeroval는 함수를 호출할 때의 값을 ival에 복사하여 가져온다.
*/
func zeroval(ival int) {
	ival = 0
}

/*
	zeroptr는 int형 포인터를 받도록 *int타입을 파라미터로 갖고 있다.
	함수안에서 *iptr는 메모리 주송레서 해당 주소의 현재 값으로 포인터를 역참조한다.
	역참조된 포인터에 값을 할당하면 이는 참조된 주소의 값을 바꾼다.
*/
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i는 i의 메모리 주소를 반환한다.
	// 즉, i의 포인터이다.
	// 포인터도 출력할 수 있다.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

/*
	zeroval은 main에 있는 i값을 바꾸지 않지만,
	zeroptr은 해당 값의 메모리 주소를 참조하고 있기 때문에
	값을 바꿀 수 있다.
*/
