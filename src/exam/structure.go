package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	// 구조체
	p := Person{}
	p.name = "Im"
	p.age = 24
	fmt.Println(p)
	fmt.Println()

	_p := new(Person)
	_p.name = "Kim"
	_p.age = 25
	fmt.Println(_p)
	fmt.Println(*_p)
	fmt.Println()

	// 생성자
	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(*dic)
	fmt.Println()
}
