package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]|)ch", "peach")
	fmt.Println(match)
	// 위애서는 문자열 패턴을 바로 사용했지만, 다른 정규식 작업을 위해선
	// 최적화된 Regexp 구조체를 Compile 해야한다.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 이 구조체에선 많은 메소드를 사용할 수 있다.
	fmt.Println(r.MatchString("peach"))

	// 정규식과 일치하는 문자열을 찾는다.
	fmt.Println(r.FindString("peach punch"))

	// 첫번째로 매칭되는 문자열을 찾지만 일치하는 텍스트
	// 대신 일치하는 텍스트의 첫 인덱스와 마지막 인덱스를 반환한다.
	fmt.Println(r.FindStringIndex("peach punch"))

	// Submatch 변형은 전체 패턴 일치와 해당 일치의 부분 일치에 대한
	// 정보를 모두 포함한다. 예를 들어 아래 코드는
	// p([a-z]+)ch와 ([a-z])에 대한 정보를 모두 반환한다.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 전체 일치와 부분 일치의 인덱스에 대한 정보를 반환한다.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// All 변형은 입력값의 첫번째만이 아닌 모든 일치 작업에 적용된다.
	// 예를 들면 정규식에 대해 모든 일치 항목들을 찾는 경우가 있다.
	fmt.Println(r.FindAllString("peach punch", -1))

	// All 변형은 위에서 본것과 마찬가지로 다른 함수에 대해 사용할 수 있다.
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 다음 함수의 두번째 인자에 음이 아닌 정수를 전달하면 일치 항목의 갯수를 제한한다.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 위의 예시들은 문자열 인자를 사용하여 MatchString과 같은 이름을 사용했다.
	// 함수명에서 String을 없애고 인자로 []byte를 전달할 수도 있다.
	fmt.Println(r.Match([]byte("peach")))

	// 정규표현식으로 상수를 만들때 Compile의 변형인 MustCompile을 사용할 수 있다.
	// 일반 Compile은 반환값이 2개라 상수로 사용할 수 없다.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// regexp 패키지는 부분분자열을 다른값으로 바꾸는데 사용할 수 있다.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 변형을 사용하여 주어진 함수로 일치돤 텍스트를 변환시킬 수 있다.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
