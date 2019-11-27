package main

import "fmt"

/*
keyword

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
*/

/*
Data Type

부울린 타입 : bool
문자열 타입 : string: string은 한번 생성되면 수정될 수 없는 Immutable 타입임
정수형 타입 : int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
Float 및 복소수 타입 : float32 float64 complex64 complex128
기타 타입 : byte: uint8과 동일하며 바이트 코드에 사용 / rune: int32과 동일하며 유니코드 코드포인트에 사용한다
*/

/*
func variable() {
	var a int = 1
	var b float32 = 11.
	var i, j, k int = 1, 2, 3
}

func constant() {
	const c int = 10
	const s string = "Hi"
}

func _for() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
  }
	println(sum)
	
	n := 1
	for n < 100 {
		n *= 2      
	}
	println(n)
	
	names := []string{"홍길동", "이순신", "강감찬"}
 
	for index, name := range names {
		println(index, name)
	}
}

func VariableParameters(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}
*/

func multiReturn(nums ...int) (int, int) {
	s := 0      // 합계
	count := 0  // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func main() {
	fmt.Println(`나는 천재이지만
저친구는 아니다.`)
	count, total := multiReturn(1, 7, 3, 5, 9)
	println(count, total) 
}