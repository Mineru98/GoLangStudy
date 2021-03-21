package main

import (
	"fmt"
	"time"
)

/*
	Timer는 미래의 한 시점에 무언가를 하고 싶을때 사용할 수 있다.
	ticker는 일정한 간격으로 무언가를 반복하자 할 때 사용할 수 있다.
	여기에 우리가 중단할 때까지 주기적으로 작업을 반복하는 티커
*/
func main() {
	/*
		티머는 타이머와 유사한 메커니즘으로 사용한다.
		값을 전송하는 채널을 사용.
		다음은 매 500ms마다 도착하는 값들을 순회하기 위해 채널에서 range를 사용한다.
	*/
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	/*
		티커는 타이머와 유사한 방식으로 중단될 수 있다.
		티커가 멈추면 티커의 채널은 더 이상 값들을 수신하지 않는다.
		다음은 티커를 1600ms 후에 중단시킨다.
	*/
	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
