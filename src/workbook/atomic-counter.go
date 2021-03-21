package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
	Go에서 상태를 관리하는 가장 기본적인 메커니즘은 채널을 통한 통신이다.
	이는 worker pool 예시를 통해 해보았다.
	Go에는 상태를 관리하기 위한 몇가지 방법이 더 있다.
	이번 예제에선 여러개의 고루틴에서 접근되는 원자성 카운터를 위한
	sync/atomic 패키지를 사용하여 살펴보겠다.
*/
func main() {
	// 카운터를 나타내기 위해 부호 없는 정수를 샤용한다.
	var ops uint64 = 0
	/*
		동시 업데이트를 시뮬레이션 하기 위해, 1 밀리초마다
		카운터를 증가시키는 고루틴을 50개 띄운다.
	*/
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 카운터를 원자적으로 증가시키기 위해 ops 카운터의 메모리 주소를
				// 전달하여 AddUint64를 사용한다.
				atomic.AddUint64(&ops, 1)
				// 증가 사이를 아주 짧게 대기한다.
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// ops가 누적될 수 있도록 1초간 대기한다.
	time.Sleep(time.Second)
	/*
		카운터가 다른 고루틴에 의해 증가되는 도중에 카운터를 안전하게 사용하기위해,
		LoadUint64를 사용하여 현재값의 복사본을 opsFinal로 가져온다.
		위에서도 보았듯이 값을 패치하기 위해 &ops의 메모리 주소를 전달한다.
	*/
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

// 프로그램을 실행하면 약 40,000번의 연산이 실행되었음을 보여준다.
