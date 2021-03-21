package main

import "fmt"

/*
	슬라이스는 크기가 가변이다.
	슬라이스는 복사할 경우 데이터를 참조한다.
*/
func main() {
	/*
		배열과 다르게, 슬라이스는 원소의 갯수가 아닌
		포함하는 원소들로만 작성된다. 길이가 0인 빈 배열을 만들기
		위해선 내장 함수 make를 사용하면 된다.
		아래 코드는 제로값으로 초기화된 길이가 3인 문자열 배열을 만든다.
	*/
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len은 슬라이스의 길이를 반환한다.
	fmt.Println("len:", len(s))

	/*
		이러한 기본 연산에 이어 슬라이스는 배열보다 풍부한 기능들을 기원한다.
		append는 새로운 값이 추가된 슬라이드를 반환하는 함수이다.
		주의할 건 새로운 슬라이스를 사용하기 위해선 append로부터 반환되는 값을 사용해야한다.
	*/
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	/*
		슬라이스는 복사도 가능하다.
		s와 동일한 길이를 갖는 빈 슬라이스 c를 생성하고 s를 c로 복사한다.
	*/
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	/*
		슬라이스는 slice[low:high]로 쓸 수있는 slice 연산을 지원한다.
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// 한줄로 슬라이스 선언 및 초기화를 할 수 있다.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	/*
		슬라이스는 다차원 데이터를 구성할 수도 있다.
		다차원 배열과 다르게 다중 슬라이스의 내부 슬라이스들은 가변적 길이를 가질 수 있다.
	*/
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	/*
		슬라이스는 배열과는 다른 타입이지만
		배열과 유사한 형태로 출력이 된다.
	*/
	fmt.Println("2d:", twoD)
}
