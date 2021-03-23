package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

/*
	파일을 만들고, 값을 쓴 다음 다 끝나면 파일을 종료시키고 싶다고 해보자.
	다음은 defer를 사용해 이를 구현한 예시이다.
*/
func main() {
	/*
		createFile로 파일 객체를 얻은 직후,
		closeFile로 파일을 종료시키는 작업을 지연한다.
		이는 writeFile이 끝나고, 함수가 끝나는 지점에서 실행된다.
	*/
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
