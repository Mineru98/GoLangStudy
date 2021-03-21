package main

import "fmt"

/*
	채널 닫기는 더 이상 채널에 보낼 데이터가 없음을 나타낸다.
	이는 채널의 리시버들에게 완료 상태를 전달하는데에 유용하다.
*/

func main() {
	/*
		고루틴에서 워커 고루틴으로 작업을 전달하기 위해 jobs 채널을 사용한다.
		워커에서 돌릴 job이 더이상 없을 경우, jobs 채널을 close 한다.
	*/
	jobs := make(chan int, 5)
	done := make(chan bool)

	/*
		아래 코드는 워커 고루틴이다.
		이는 j, more := <- jobs를 통해 jobs로부터 반복적으로 값을 수신한다.
		이 두값을 반환하는 특별한 형태의 수신값에서, more 값은 jobs가
		close 되고 채널에 있는 모든 값들이 수신될 경우 false 값을 갖게 된다.
		모든 job이 종료되었음을 알리기 위해 done을 사용한다.
	*/
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// jobs 채널을 통해 워커로 3개의 job을 보낸 후, 채널을 닫는다.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// job이 끝날때 까지 done을 수신받으면서 고루틴 작업이 끝날때까지 대기를 한다.
	<-done
}
