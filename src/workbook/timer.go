package main

import (
	"fmt"
	"time"
)

/*
	종종 Go 코드를 나중에 특정 시점에서 실행하거나 일정 간격으로
	반복 실행시키고 싶을 때가 있다.
	Go의 내장 기능인 timer와 ticker는 이를 쉽게 만들어준다.
*/

func main() {
	/*
		타이머는 미래의 한 이벤트를 나타낸다.
		원하는 대기 시간만큼을 타이머에게 지정해주면,
		이는 해당 시각에 알림을 주는 채널을 반환한다.
		다음 타이머는 2초를 대기한다.
	*/
	timer1 := time.NewTimer(time.Second * 2)
	/*
		<-timer1.C는 타이머가 만료되었음을 알려주는 값을 보내기 전까지
		타이머의 C 채널을 블로킹한다.
	*/
	<-timer1.C
	fmt.Println("Timer 1 expired")

	/*
		그저 대기만 하고자 했을땐, time.Sleep을 사용할 수 있었다.
		타이머가 유용한 이유 중 하나는 타이머가 만료되기 전에 취소 시킬 수 있다는 것이다
	*/
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

/*
	첫번째 타이머는 프로그램을 시작한지 약 2초 뒤에 만료가 되지만,
	두번째는 만료가 되기 전에 중단되어야 한다.
*/
