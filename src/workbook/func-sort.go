package main

/*
	가끔 일반적인 순서가 아닌 특정한 순서로 컬렉션을 정렬하고 싶을 때가 있다.
	예를 들면, 문자열들을 알파벳 순서가 아닌 길이를 기준으로 정렬하고 싶은 경우가 있다.
*/
import (
	"fmt"
	"sort"
)

/*
	Go에서 커스텀 함수로 정렬을 하려면, 그에 해당하는 타입이 필요하다.
	우리는 여기서 내장 타입인 []string의 alias인 ByLength라는 타입을 만들었다.
*/

type ByLength []string

/*
	여기서 정의한 타입에 sort.Interface-Len, Less, Swap를 구현하면
	sort 패키지의 제네릭 Sort 함수를 사용할 수 있다.
	Len과 Swap은 일반적으로 타입에 따라 유사하며 Less가 실제 커스텀 정렬 로직을 갖는다.
	우리의 경우 문자열 길이가 증가하는 순서로 정렬을 하고자 하므로,
	len(s[i])와 len(s[j])를 사용한다.
*/
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

/*
	구현한 것들은 가지고 원래의 fruits 슬라이스를 ByLength로 캐스팅하고
	여기에 sort.Sort를 사용함으로써 우리는 커스텀 정렬을 구현할 수 있다.
*/
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

/*
	프로그램을 실행하면 우리가 원하던대로 문자열 길이를 기준으로 정렬된 리스트가 보여진다.

	커스텀 타입을 생성하고, 해당 타입에 세 개의 Interface 메소드를 구현한 다음
	커스텀 타입의 컬렉션에 sort.Sort를 호출하는 패턴을 따르면 어떤 함수로도
	Go 슬라이스를 정렬할 수 있다.
*/
