package main

import (
	"fmt"
	"time"
)

func sendChan(ch chan <- string) {
	ch <- "Data"
}

func receiveChan(ch <- chan string) {
	data := <- ch
	fmt.Println(data)
}

func run1(done chan bool) {
	time.Sleep(time.Second * 1)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(time.Second * 2)
	done <- true
}

func main() {
	/*
		Go 채널은 채널을 통하여 데이터를 주고 받는 통로라고 볼 수 있다.
		채널은 make() 함수를 통해 미리 생성되어야 하며, 채널 연산자 <- 를 통해 데이터를 보내고 받는다.
		채널은 흔히 Go루틴들 사이에서 데이터를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서
		대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용 된다.
	*/
	// 정수형 채널 생성
	ch := make(chan int)
	go func() {
		ch <- 123
	}()
	var i int
	i = <- ch
	fmt.Println(i)
	fmt.Println()
	/*
		Go 채널은 수신자와 송신자가 서로 기다리는 속성때문에, 이를 이용하여
		Go 루틴이 끝날 때까지 기다리는 기능을 구현할 수 있다.
		즉, 익명함수를 사용한 한 Go 루틴에서 어떤 작어비 실행 되고 있을 때,
		메인루틴은 <- done에서 계속 수신하며 대기하고 있게 된다.
		익명 함수 Go루틴에서 작업이 끝날 후, done채널에 true를 보내면,
		수신자 메인루틴은 이를 받고 프로그램을 종료하게 됨
	*/
	done := make(chan bool)
	go func() {
		for i:= 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	
	result := <- done
	fmt.Println(result)
	fmt.Println()
	
	// Go 채널 버퍼링
	/*
		Go 채널은 2가지 채널이 있다.
		Unbuffered channel, Buffered channel
		위의 두가지 예시는 Unbuffered channel로서 이 채널에서는 하나의 수신자가 데이터를 받을 때까지
		송신자가 데이터를 보내는 처널에 묶여 있게 된다.
		하지만, Buffered channel을 사용하면 비록 수신자가 받을 준비가 되어 있지 않을지라도
		지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있다.
		버퍼 채널은 make(chan type, N) 함수를 통해 생성되는데, 두번재 파라미터 N에 사용할 버퍼 갯수를 넣는다.
	*/
	// c := make(chan int)
	// c <- 1 // 수신루틴이 없으므로 데드락
	c := make(chan int, 1)
	c <- 101
	fmt.Println(<-c)
	fmt.Println()
	
	// 채널 파라미터
	/*
		채널을 함수의 파라미터도 전달할 때, 일반적으로 송수신을 모두 하는 채널을 전달하지만,
		특별히 해당 채널로 송신만 할 것인지 혹은 수신만할 것인지를 지정할 수도 있다.
		송신 파라미터는 (p chan <- int) 와 같이 chan<- 을 사용하고,
		수신 파라미터는 (p <- chan int) 와 같이 <-chan 을 사용한다.
		만약 송신 채널 파라미터에서 수신을 한다거나, 수신 채널에 송신을 하게 되면, 에러가 발생한다.
	*/
	
	chex := make(chan string, 1)
	sendChan(chex)
	receiveChan(chex)
	fmt.Println()
	
	// 채널 닫기
	chClose := make(chan int, 2)
	chClose <- 1
	chClose <- 2
	
	close(chClose)
	
	fmt.Println(<-chClose)
	fmt.Println(<-chClose)
	if _, success := <-chClose; !success {
		println("더이상 데이타 없음.")
	}
	fmt.Println()
	
	// 채널 range문
	chRan := make(chan int, 2)
	
	chRan <- 1
	chRan <- 2
	
	close(chRan)
	
	// for {
	// 	if i, success := <-chRan; success {
	// 		fmt.Println(i)
	// 	} else {
	// 		break
	// 	}
	// }
	
	for i := range chRan {
		fmt.Println(i)
	}
	fmt.Println()
	
	// 채널 select문
	done1 := make(chan bool)
	done2 := make(chan bool)
	
	go run1(done1)
	go run2(done2)
EXIT:
	for {
		select {
		case <- done1:
			fmt.Println("run1 완료")
		case <- done2:
			fmt.Println("run2 완료")
			break EXIT
		}
	}
}