package main

// sort 패키지는 내장 타입과 사용자 정의 타입을 위한 정렬을 구현하고 있다.

import (
	"fmt"
	"sort"
)

func main() {
	/*
		정렬 method는 내장 타입에만 해당한다. 다음은 문자열에 대한 예시이다.
		참고로 정렬은 제자리 정렬이므로,
		이는 전달된 슬라이스를 변형시키며 새로운 슬라이스를 반환하진 않는다.
	*/
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("------------------")
	// int형의 정렬 예시이다.
	ints := []int{7, 2, 4}
	s := sort.IntsAreSorted(ints)

	fmt.Println("Ints:", ints)
	fmt.Println("Sorted:", s)

	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	/*
		또한 sort를 사용해 슬라이스가 이미 정렬되어 있는지에
		대한 여부도 확인할 수 있다.
	*/
	s = sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}

/*
	프로그램을 실행하면 정렬된 문자열과 정수 슬라이스
	그리고 AreSorted 테스트의 결과로 true로 출력된다.
*/
