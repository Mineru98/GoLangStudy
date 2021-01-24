package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 다중 CPU처리(멀티 코어 프로세싱)
	runtime.GOMAXPROCS(runtime.NumCPU()) // 현재 CPU의 모든 코어 개수만큼 사용
	
	// Go루틴
	say("Sync") // 함수를 동기적으로 실행
	
	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")
	
	time.Sleep(time.Second * 1) // 1초 대기
	fmt.Println()
	
	// 익명함수 Go루틴
	// WaitGorup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)
	
	go func() {
		defer wait.Done()
		time.Sleep(time.Second * 1) // 1초 대기
		fmt.Println("World")
	}()
	
	go func(msg string) {
		defer wait.Done()
		time.Sleep(time.Second * 2) // 2초 대기
		fmt.Println(msg)
	}("Hello")
	
	wait.Wait()// Go루틴이 모두 끝날때까지 대기
	fmt.Println()
}