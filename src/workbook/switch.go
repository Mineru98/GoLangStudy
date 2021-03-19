package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as")

	switch i {
	case 1:
		fmt.Printf("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*
		동일한 case문에서 여러개의 표현식을 구분하기 위해
		콤마를 사용할 수 있다. 이 예시에서 default를 사용했다.
		default는 선택 사항이다.
		개인적으로 다르다고 생각한 부분은 break를 굳이 쓰지 않는다는 것이다.
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/*
		표현식이 없는 switch는 if/else 로직을 표현하기
		위한 또 다른 방법이다. 여기서 상수가 아니 case문을
		사용하는 방법을 볼 수 있다.
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	/*
		타입 switch는 값 대신 타입을 비교한다.
		인터페이스 값의 타입을 알아내기 위해 이를 사용할 수 있다.
		이 예시에서는, 변수 t는 해당 절에 해당하는 타입을 가질 것이다.
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
