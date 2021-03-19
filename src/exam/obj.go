package main

import (
	"fmt"
)

// Rect - struct 정의
type Rect struct {
	width, height int
}

// Rect의 area() 메소드
func (r Rect) area() int {
	r.width++
	return r.width * r.height
}

// Pointer Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	// 객체 메소드
	rect := Rect{9, 20}
	area := rect.area()
	fmt.Println(rect)
	fmt.Println(area)
	fmt.Println()

	// Value vs Pointer
	/*
		Value receiver는 구조체의 데이터를 복사하여 전달
		Pointer receiver는 구조체의 포인터만 전달

		Value receiver의 경우, 메소드 안에서 해당 구조체의 필드값이 변경되더라도 호출자의 데이터는 변경 되지 않는다.
		Pointer receiver의 경우, 메소드 내의 필드 값이 변경이 그대로 호출자에서 반영이 된다.
	*/
	rect2 := Rect{9, 20}
	area2 := rect2.area2()
	fmt.Println(rect2)
	fmt.Println(area2)
	fmt.Println()
}
