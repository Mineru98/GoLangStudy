package main

import (
	"fmt"
	"log"
	"os"
)

type error interface {
	Error() string
}

func otherFunc(filePath string) (f string, err string) {
	f, err = os.Open(filePath)
	return
}

func main() {
	// Error 처리
	f, err := os.Open("src/example/main.go")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(f.Name())
	fmt.Println()

	// 활용 방법은 좀 더 알아야 할 듯...
}
