package main

import (
	"fmt"
	"os"
)

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func openFile2(fn string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Open Error!")
		}
	}()
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func main() {
	// 지연 실행 defer
	f, err := os.Open("./src/example/exam.txt")
	if err != nil {
		panic(err)
	}

	// main 함수 가장 마지막에 f.close가 실행 됨
	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	fmt.Println(len(bytes))
	fmt.Println()

	// panic 함수
	/*
		현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 반환함.
		이러한 panic 모드 실행 방식은 다시 상위 함수에도 똑같이 적용 되고,
		계속 콜스택을 타고 올라가며 적용 된다. 그리고 마지막에는 프로그램이 에러를 내고 종료하게 됨.
	*/

	// 	openFile("Invalid.txt")
	// 	fmt.Println("Done")
	// 	fmt.Println()

	// recover 함수
	/*
		panic 함수에 의한 패닉상태를 다시 정상 상태로 되돌리는 함수.
		위의 panic 함수 로직에서는 마지막 Done을 출력하지 못하고 프로그램이 crash 되지만,
		recover 함수를 사용하면 panic 상태를 제거하고 나머지 문장들을 실행이 가능하다.
	*/
	openFile2("Invalid.txt")
	fmt.Println("Done")
	fmt.Println()
}
