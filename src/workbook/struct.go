package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// 새로운 구조체를 생성한다.
	fmt.Println(person{"Bob", 20})
	//필드명을 사용해 구조체를 초기화 할 수 있다.
	fmt.Println(person{name: "Alice", age: 30})
	// 생략 된 필드는 제로값을 갖게 된다.
	fmt.Println(person{name: "Fred"})
	// 앞에 &를 붙이면 구조체 포인터를 생성할 수 있다.
	fmt.Println(&person{name: "Ann", age: 40})
	// . 을 사용해 구조체 필드에 접근한다.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 구조체 포인터에서도 . 을 사용할 수 있다.
	// 이때 포인터는 자동으로 역참조 된다.
	sp := &s
	fmt.Println(sp.age)
	// 구조체는 수정이 가능(mutable)하다.
	sp.age = 51
	fmt.Println(sp.age)
}
