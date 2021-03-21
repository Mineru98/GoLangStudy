package main

import "fmt"

type rect struct {
	width, height int
}

// area 메소드는 *react를 리시버 타입으로 갖는다.
func (r *rect) area() int {
	return r.width * r.height
}

/*
	메소드는 포인터 또는 값을 리시버 타입으로하여 정의될 수 있다.
	아래 코드는 값을 리시버로 한다.
*/
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	// 구조체에 정의된 두 개의 메소드를 호출했다.
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	/*
		Go는 메소드 호출에 대해 값과 포인터간의 변환을 자동으로 처리한다.
		메소드 호출 시 값 복사를 피하기 위해 포인터 리시버 타입을 사용할 수도 있고
		전달되는 구조체를 변경할 수 있도록 할 수도 있다.
	*/
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
