package main

/*
	timeout은 외부 리소스를 연결하거나 실행 시간을 제한해야하는 프로그램에서
	중요하다. Go에서 타임아웃 구현은 채널과 select를 사용하면 쉽고 우아하게 할 수 있다.
*/

import (
	"fmt"
	"time"
)

func main() {
	// 2초 후에 c1 채널에 결과값을 반환하는 외부 호출을 실행한다고 가정하자.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	/*
		다음은 select로 구현한 타임아웃이다.
		res := <-c1은 결과값을 대기하고 <-Time.After는 1초의 타임아웃 후에
		전달되는 값을 대기한다. select는 대기중인 첫번째 리시브로 진행되어 있기 때문에,
		만약 이 연산이 허용된 1초 보다 더 오래걸릴 경우 타임아웃이 발생한다.
	*/
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 만약 타임아웃을 3초로 늘리면, c2로부터의 수신은 성공하고 결과값을 출력할것이다.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

/*
	이 프로그램을 실행하면 첫 연산은 타임아웃되고
	두번째 연산은 성공하는걸 볼 수 있다.

	select 타임아웃 패턴을 사용하려면 채널을 통해 결과값을 전달해야한다.
	Go의 중요한 기능들이 채널과 select 기반이기 때문에 일반적으로
	이는 좋은 아이디어이다.
*/
