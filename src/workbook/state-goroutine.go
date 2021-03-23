package main

/*
	채널 기반 접근법은 Go의 통신을 통한 메모리 공유와
	정확히 한 고루틴이 각 데이터의 일부를 소유한다는 아이디어에 기반한다.
*/

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

/*
	상태는 한 고루틴에 의해 소유된다.
	이는 데이터가 동시 접근으로인해 손상되지 않음을 보장한다.
	상태를 읽거나 쓰기위해, 다른 고루틴들은 소유중인 고루틴으로
	메시지를 보내고 응답을 받는다.
	다음 readOps와 writeOps 구조체는 요청들과 소유중인 고루틴이
	응답하기 위한 방법을 캡슐화 한다.
*/

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 이전과 같이 얼마나 연산이 얼마나 이루어지는지를 카운팅한다.
	var readOps uint64 = 0
	var writeOps uint64 = 0
	/*
		reads와 writes 채널은 각각 읽기와 쓰기 요청을 전달하기
		위함으로 다른 고루틴들에 의해 사용된다.
	*/
	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	/*
		다음 고루틴은 state를 소유하며, 이는 이전 예제에서처럼 map을
		사용하지만 지금은 상태있는 고루틴이 이를 프라이빗으로 갖고 있다.
		이 고루틴은 반복적으로 reads와 writes 채널중 하나를 선택하며,
		값이 도착하는대로 요청에 대한 응답 처리를 한다.
		응답은 우선 요청된 연산을 수행함으로써 실행되며 write의 경우는 성공을
		알리기 위해(reads의 경우는 원하는 값을) 응답 채널 resp로 값을 전달한다.
	*/
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	/*
		다음 reads 채널을 통해 상태를 가진 고루틴에 읽기를
		요청하기 위해 100개의 고루틴을 띄운다.
		각 읽기는 readOp을 생성해야하며, 이를 reads 채널을 통해
		전달하고, resp 채널로 들어온 값을 통해 결과값을 받는다.
	*/
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 마찬가지로 10개의 쓰기 고루틴을 띄운다.
	for w := 0; w < 100; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 1초간 고루틴들의 작업을 대기한다.
	time.Sleep(time.Second)
	// 마지막으로 각 연산 횟수를 캡쳐하고 리포팅한다.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}

/*
	이 특별한 케이스인 고루틴 기반 접근법은 뮤텍스 기반보다 조금 더 복잡하다.
	이는 특정한 경우에 유용할 수 있는데, 예를 들면 다른 채널들이 관련되어
	있거나 에러가 발생하기 쉬운 다중 뮤텍스들을 관리하는 경우가 있다.
	프로그램의 정확성을 이해하는 것과 관련해 가장 자연스럽게 느껴지는 방식을 사용해야 한다.
*/
