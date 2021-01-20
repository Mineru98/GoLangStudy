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

type person struct {
	name string
	age  int
}

type Rect struct {
    width, height int
}

func variable() {
	var a int = 1
	var b float32 = 11.
	var i, j, k int = 1, 2, 3
	fmt.Println(a,b,i,j,k)
}

func constant() {
	const c int = 10
	const s string = "Hi"
	fmt.Println(c,s)
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

func multiReturn(nums ...int) (int, int) {
	s := 0      // 합계
	count := 0  // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func closer() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

//Rect의 area() 메소드
func (r Rect) area() int {
	r.width++
	return r.width * r.height   
}

// 포인터 Receiver
func (r *Rect) area2() int {
	r.width++
	return r.width * r.height
}

func main() {
	println("***변수***")
	variable()
	println("***상수***")
	constant()
	println("***for문***")
	_for()
	println("***가변 함수***")
	VariableParameters("Test","Message")
	println("***다중 리턴 함수***")
	{
		count, total := multiReturn(1, 7, 3, 5, 9)
		fmt.Println(count, total)
	}
	println("***익명 함수***")
	{
		sum := func(n ...int) int { //익명함수 정의
			s := 0
			for _, i := range n {
				s += i
			}
			return s
		}

		result := sum(1, 2, 3, 4, 5) //익명함수 호출
		println(result)
	}
	println("***익명 함수의 활용***")
	{
		//변수 add 에 익명함수 할당
		add := func(i int, j int) int {
			return i + j
		}
		// add 함수 전달
		r1 := calc(add, 10, 20)
		println(r1)
		// 직접 첫번째 파라미터에 익명함수를 정의함
		r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
		println(r2)
	}
	println("***클로저***")
	{
		next := closer()
		println(next())  // 1
		println(next())  // 2
		println(next())  // 3
		anotherNext := closer()
		println(anotherNext()) // 1 다시 시작
		println(anotherNext()) // 2
		println(next())  // 4
	}
	println("***컬렉션***")
	{
		var a [3]int  //정수형 3개 요소를 갖는 배열 a 선언
    a[0] = 100
    a[1] = 200
    a[2] = 300
		for index, num := range a {
			println(index, num)
		}
		
		// var a1 = [3]int{1, 2, 3}
		// var a3 = [...]int{1, 2, 3} // 가변배열
	}
	println("***컬렉션 Slice***")
	{
		s := make([]int, 5, 10) // make는 Go의 내장 함수이다.
    println(len(s), cap(s)) // len은 현재 할당 된 배열의 크기이고, cap은 배열의 최대 길이를 지정한 크기 값이다.
		for index, num := range s {
			println(index, num)
		}
		s = append(s, 3, 4, 5) // 배열 추가
		for index, num := range s {
			println(index, num)
		}
		
		sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
    sliceA = append(sliceA, sliceB...) //sliceA = append(sliceA, 4, 5, 6)
    fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
	}
	println("***컬렉션 Map***")
	{
		// var idMap map[int]string
		tickers := map[string]string{
			"GOOG": "Google Inc",
			"MSFT": "Microsoft",
			"FB":   "FaceBook",
			"AMZN": "Amazon",
		}
		
		// map 키 체크
    val, exists := tickers["MSFT1"]
    if !exists {
        println("No MSFT ticker")
		} else {
			println(val)
		}
		
		for key, val := range tickers {
			println(key, val)
		}
		
		var m map[int]string
 
    m = make(map[int]string)
    //추가 혹은 갱신
    m[901] = "Apple"
    m[134] = "Grape"
    m[777] = "Tomato"
 
    // 키에 대한 값 읽기
    str := m[134]
    println(str)
 
    noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
    println(noData)
 
    // 삭제
    delete(m, 777)
		for key, val := range m {
			println(key, val)
		}
	}
	println("***객체***")
	{
		p := person{}
    // 필드값 설정
    p.name = "Lee"
    p.age = 10
     
    fmt.Println(p)
		
		// 객체 할당 방법 두가지
		var p1 person 
		p1 = person{"Bob", 20}
		p2 := person{name: "Sean", age: 50}
		fmt.Println(p1, p2)
	}
	println("***객체 메서드***")
	{
		rect := Rect{10, 20}
    area := rect.area() //메서드 호출
    println(rect.width, area) // 10 220 출력
		
		rect2 := Rect{10, 20}
    area2 := rect2.area2() //메서드 호출
    println(rect2.width, area2) // 11 220 출력
	}
	/*http://golang.site/go/article/18-Go-%EC%9D%B8%ED%84%B0%ED%8E%98%EC%9D%B4%EC%8A%A4 부터 공부*/
	{
		
	}
}