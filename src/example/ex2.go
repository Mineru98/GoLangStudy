package main

import (
	"fmt"
)

func main() {
	// 데이터 타입
	/*
		부울린 타입
		bool
		
		문자열 타입
		string(immuatble 타입)
		
		정수형 타입
		int int8 int16 int32 int64
		unit unit8 unit16 unit32 uint64 unitptr

		복소수 타입
		float32 float64 complex64 complex128

		기타 타입
		byte(unit8과 동일)
		rune(int32와 동일)
	*/
	
	// 문자열
	rawLiteral := `아리랑\n아리랑\n아라리요`
	interLiteral := "아리랑아리랑\n아라리요"
	
	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)
	
	// 데이터 타입 캐스팅
	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	fmt.Println(u, f)
	
	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	fmt.Println(bytes, str2)
}