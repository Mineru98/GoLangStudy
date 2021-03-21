package main

import "fmt"

/*
	채널의 송수신은 기본적으로 동기적이다.
	그러나, select를 default문과 함께 사용하면, 비동기 송수신을
	구현할 수 있으며, 비동기 다중 select도 구현이 가능하다.
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	/*
		아래 코드는 비동기 수신이다.
		messages에서 값을 사용할 수 있는 경우,
		select는 <-message case에서 그 값을 가져온다.
		그렇지 않은 경우엔 바로 default case를 수행한다.
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message recevied")
	}
	// 비동기 송신도 유사하게 동작한다.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("no message sent")
	}
	/*
		다중 비동기 select를 구현하기 위해 위의 default 문에 다중 case를 구현할 수 있다.
		다음은 messages와 signals 두 채널로부터의 비동기 수신을 시도한다.
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
