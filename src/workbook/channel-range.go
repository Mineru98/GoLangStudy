package main

import "fmt"

// for와 range가 기본 자료구조에서 어떻게 순회하는지 살펴보자.
func main() {
	// queue 채널에 있는 2개의 값을 순회 해보겠다.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	/*
		이 range는 queue로부터 수신된 각 원소들을 순회한다.
		위에서 채널을 종료했기 때문에 이 순회 루프는 2개의 원소를 받으면 종료된다.
	*/
	for elem := range queue {
		fmt.Println(elem)
	}
}
