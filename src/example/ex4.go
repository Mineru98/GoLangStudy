package main

import (
	"fmt"
)

func grade(score int) {
	switch {
	case score >= 90:
		fmt.Println("A\n")
	case score >= 80:
		fmt.Println("B\n")
	case score >= 70:
		fmt.Println("C\n")
	case score >= 60:
		fmt.Println("D\n")
	default:
		fmt.Println("F\n")
	}
}

func main() {
	// 조건문
	var k int = 1
	if k == 1 {
		fmt.Println("One\n")
	}
	
	k = 2
	if k == 1 {
		fmt.Println("One\n")
	} else if k == 2 {
		fmt.Println("Two\n")
	} else {
		fmt.Println("Other\n")
	}
	
	max := 10
	// val에 k * 2를 할당하고 바로 이어서 if 문을 실행하는 문장
	if val := k * 2; val < max {
		fmt.Println(val)
	}
		fmt.Println()
	
	// 스위치 문
	/*
		1. switch 뒤에 expression이 없을 수 있음. : 다른 언어는 switch 키워드 뒤에 변수나
																						expression 반드시 두지만, Go는 이를 쓰지 않아도 된다.
																						이 경우는 Go는 switch expression을 true로 생각하고
																						첫번째 case문으로 이동하여 검사한다.
		2. case 문에 expression을 쓸 수 있음. : 	다른 언어의 case문의 일반적으로 리터럴 값만을 갖지만, 
																					Go는 case문에 복잡한 expression을 가질 수 있다.
		3. No default fall through :  다른 언어의 case문은 break를 쓰지 않는 한 다음 case로 이동하지만,
																	Go는 다음 case로 가지 않는다.
		4. Type switch :  다른 언어의 switch는 일반적으로 변수의 값을 기준으로 case로 분기하지만,
											Go는 그 변수의 Type에 따라 case로 분기할 수 있다.
	*/
	var name string
	var category = 1
	
	switch category {
	case 1:
		name = "Paper Book\n"
	case 2:
		name = "eBook\n"
	case 3, 4:
		name = "Blog\n"
	default:
		name = "Other\n"
	}
	fmt.Println(name)

	switch x:= category << 2; x - 1 {
	case 1:
		name = "Paper Book\n"
		fallthrough
	case 2:
		name = "eBook\n"
		fallthrough
	case 3, 4:
		name = "Blog\n"
		fallthrough
	default:
		name = "Other\n"
	}
	fmt.Println(name)
	
	grade(95)
	grade(85)
	grade(65)
	grade(55)
	grade(45)
	
}