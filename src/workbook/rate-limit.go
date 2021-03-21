package main

/*
	속도제한은 리소스 이용을 제어하고 서비스의 품질을 유지하기 위한
	중요한 메커니즘이다. Go는 고루틴, 채널 그리고 티커로 속도 제한을 지원한다.
*/

import (
	"fmt"
	"time"
)

func main() {
	/*
		요청 핸들링을 제한하고 싶다고 가정해보자.
		requests라는 이름으로 요청을 처리하고 채널을 만든다.
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// limiter 채널은 매 200 밀리초마다 값을 받는다.
	// 이게 여기서 사용할 속도 제한 조절기이다.
	limiter := time.Tick(time.Millisecond * 200)
	/*
		각 요청을 처리하기 전에 limiter 채널의 수신으로 블로킹함으로써
		매 200 밀리초마다 하나의 요청을 받도록 제한하고 있다.
	*/
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	/*
		전반적으로는 속도 제한을 유지하면서 요청들을 짧게 버스팅하고 싶은 경우가 있다.
		limiter 채널을 버퍼링함으로써 이를 가능케할 수 있다.
		burstyLimiter 채널로 최대 3개의 이벤트를 버스팅할 수 있다.
	*/
	burstyLimiter := make(chan time.Time, 3)
	// 허용된 버스팅 수만큼 채널을 채운다.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 매 200 밀리초마다 최대 3개까지 burstyLimiter에 새로운 값을 추가한다.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	/*
		5개의 요청이 들어오는 시뮬레이션을 해보자.
		처음 3개의 요청은 burstyLimiter의 버스팅 기능의 이점을 얻는다.
	*/
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*
	프로그램을 실행하면 첫번째 요청 배치들이 우리가 원하는대로
	매 약 200 밀리초마다 처리되고 있음을 볼 수 있다.

	두번째 요청 배치에선 첫 3개의 요청은 버스팅 가능한 속도 제한덕에
	즉각적으로 처리되며 나머지 2개의 요청은 각각 약 200 밀리초의 딜레이를 가지고 처리된다.
*/
