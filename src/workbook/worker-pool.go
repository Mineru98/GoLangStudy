package main

import (
	"fmt"
	"time"
)

/*
	여러개의 인스턴스를 동시에 실행시킬 워커이다.
	이 워커들은 jobs 채널을 통해 작업을 받으며 작업의 결과값을
	results로 보낸다. 비용이 큰 작업을 시뮬레이션 하기 위해
	각 job마다 1초의 딜레이를 줄것이다.
*/
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// 워커풀을 사용하기 위해선 워커에 작업을 보내고 그 결과값들을 수집해야한다.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 다음은 3개의 워커를 실행시키는데,
	// 처음에는 job이 없기 때문에 각 워커는 블로킹 된다.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 5개의 job을 보내고 작업을 다 보냈음을 알리기위해 채널을 close 한다.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 마지막으로 모든 작업의 결과값들을 가져온다.
	for a := 1; a <= 5; a++ {
		<-results
	}
}

/*
	프로그램은 5개의 job이 여러개의 워커에 의해 실행되고
	있음을 보여준다. 총 작업 시간은 5초지만 3개의 워커가
	동시에 실행되고 있기 때문에, 전체 작업은 약 2초만에 끝난다.
*/
