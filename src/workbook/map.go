package main

import "fmt"

/*
	Go에서 map은 다른 언어에서는
	hashes나 dicts라고 부른다.
*/

func main() {
	// 빈 map을 생성하기 위해 내장 함수인 make를 사용한다.
	// make(map[key-type]value-type)
	m := make(map[string]int)

	// name[key] = val로 key/value 쌍을 저장한다.
	m["k1"] = 7
	m["k2"] = 13

	// map 출력
	// map은 모든 key/value 쌍들을 출력한다.
	fmt.Println("map:", m)

	// name[key]로 키에 해당하는 값을 가져온다.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len을 map에 사용하면, key/value 쌍의 갯수를 반환한다.
	fmt.Println("len:", len(m))

	// delete 내장 함수는 map의 key/value 쌍을 key로 삭제한다.
	delete(m, "k2")
	fmt.Println("map:", m)

	/*
		map에서 값을 꺼내오면서 반환되는 선택적인 두번째 반환값은 해당 key가 map에
		존재하는지에 대한 여부를 나타낸다.
		이는 key 값이 존재하지 않는건지 또는 해당 값이 0이나 ""과 같은 제로값인지를
		명확하게 구분하는데 사용할 수 있다.
		값 자체가 필요없는 경우엔, 공백 식별자 _를 사용해 무시할 수 있다.
	*/
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	// 한줄로 map 선언 및 초기화가 가능하다.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
