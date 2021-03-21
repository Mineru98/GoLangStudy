package main

/*
	고루틴간의 실행을 동기화하기 위해 채널을 사용할 수 있다.
	고루틴이 끌날때까지 대기하기 위해 블로킹 리시브를 사용하는 예제이다.
*/

import (
	"fmt"
	"time"
)

type r_state struct {
	done bool
	name string
}

/*
	고루틴에서 실행하기 위한 함수이다.
	done 채널은 이 함수의 작업이 끝났음을 다른 고루틴에게
	알리기 위해 사용된다.
*/
func worker(state chan r_state, name string) {
	fmt.Println(name, "working...")
	time.Sleep(time.Second)
	fmt.Println(name, "done")
	// 작업의 종료를 알리기 위해 값을 전달한다.
	var s r_state
	s.done = true
	s.name = name
	state <- s
}

func main() {
	work_num := 5
	stack := 0
	done := make(chan r_state, work_num)
	// 알림을 위한 채널을 전달하면서 워커 고루틴을 채널 버퍼 갯수만큼 실행한다.
	go worker(done, "#1")
	go worker(done, "#2")
	go worker(done, "#3")
	go worker(done, "#4")
	go worker(done, "#5")

	for i := 0; i < work_num; i++ {
		checkSum := <-done
		if checkSum.done == true {
			stack += 1
		}
	}

	if stack == work_num {
		fmt.Println("done")
	}
}
