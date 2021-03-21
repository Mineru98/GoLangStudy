package main

import "fmt"

/*
	채널은 동시에 실행되고 있는 고루틴을 연결해주는 파이프이다.
	한 고루틴으로부터 채널에 값을 전달하면 다른 고루틴에서 이 값을 받을 수 있다.
*/
func main() {
	/*
		make(chan val-type)을 사용해 새로운 채널을 생성한다.
		채널은 전송하는 값의 타입을 타입으로 갖는다.
	*/
	messages := make(chan string)

	/*
		channel <- 구문을 사용하여 값을 채널에 전달한다.(send)
		다음은 새로운 고루틴에서
		위 코드에서 만든 messages 채널에 "ping"을 보낸다.
	*/
	go func() { messages <- "ping" }()

	/*
		<-channel 구문은 채널로부터 값을 받는다.(receives)
		다음은 위에서 전달된 "ping" 메시지를 받아 출력한다.
	*/
	msg := <-messages
	fmt.Println(msg)
}

/*
	프로그램을 실행하면 "ping" 메시지가 채널을 통해
	한 고루틴에서 다른 고루틴으로 성공적으로 전달된다.

	기본적으로 송신과 수신은 송신자와 수신자가 준비될 때까지
	블로킹된다. 이 특징은 다른 동기화 작업을 하지 않고도
	"ping" 메시지를 위해 프로그램의 종료를 기다리게 할 수 있다.
*/
