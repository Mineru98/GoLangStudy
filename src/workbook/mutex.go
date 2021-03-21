package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*
	원자성 카운터보다 복잡한 상태에 대해서는
	여러개의 고루틴이 데이터에 안전하게 접근할 수 있게
	뮤텍스를 사용할 수 있다.
*/
func main() {
	var state = make(map[int]int)
	// mutex는 state에 대한 접근을 동기화한다.
	var mutex = &sync.Mutex{}
	// 읽기와 쓰기가 얼마나 이루어지는지를 추적한다.
	var readOps uint64 = 0
	var writeOps uint64 = 0
	// 100개의 고루틴을 띄워 각 고루틴에서 1밀리초마다 반복적으로 읽기를 실행한다.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				/*
					각 읽기마다 접근을 위한 키값을 선택한다.
					state에 상호 배제 접근을 보장하기 위해 mutex를
					Lock() 한 뒤 키값으로 값을 읽고, 뮤텍스를 Unlock()
					한 다음 readOps 카운트값을 증가시킨다.
				*/
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 읽기에서와 같은 패턴으로 쓰기를 시뮬레이션 하기 위해 10개의 고루틴을 띄운다.
	for w := 0; w < 10; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}
	// 1초간 state와 mutex에서 10개의 고루틴 작업을 돌린다.
	time.Sleep(time.Second)
	// 최종 연산 횟수를 리포팅한다.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	// state의 최종 잠금으로, 어떻게 끝났는지 보여준다.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
