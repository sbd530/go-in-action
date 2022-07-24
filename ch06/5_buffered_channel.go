/*
Unbuffered channel은 채널을 통한 송수신에 무조건 블로킹이 걸린다.
하지만 Buffered channel의 채널 내의 잠금 방식은 채널의 버퍼가 꽉 찼을 때만 이루어진다.
따라서 값을 보내고 받는 동작이 동시에 이루지는 것을 보장하지 않는다.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 실행할 고루틴의 개수
	taskLoad         = 10 // 처리할 작업의 개수
)

var wg sync.WaitGroup

func init() {
	// 랜덤값 생성기 초기화
	rand.Seed(time.Now().Unix())
}

func main() {
	// 작업 부하를 관리하기 위한 버퍼있는 채널
	tasks := make(chan string, taskLoad)

	// 작업을 처리할 고루틴들을 실행
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}
	// 실행할 작업을 추가
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업: %d", post)
	}
	// 작업을 모두 처리하면 채널을 닫는다.
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 작업 할당 대기
		task, ok := <-tasks
		if !ok {
			fmt.Printf("# worker: %d : 작업 종료!\n", worker)
			return
		}

		// 작업 시작
		fmt.Printf("# worker: %d : 작업 시작: %s\n", worker, task)
		// 작업 처리 흉내
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// 작업 완료
		fmt.Printf("# worker: %d : 작업 완료: %s\n", worker, task)
	}
}

// # worker: 4 : 작업 시작: 작업: 1
// # worker: 2 : 작업 시작: 작업: 2
// # worker: 1 : 작업 시작: 작업: 4
// # worker: 3 : 작업 시작: 작업: 3
// # worker: 2 : 작업 완료: 작업: 2
// # worker: 2 : 작업 시작: 작업: 5
// # worker: 2 : 작업 완료: 작업: 5
// # worker: 2 : 작업 시작: 작업: 6
// # worker: 4 : 작업 완료: 작업: 1
// # worker: 4 : 작업 시작: 작업: 7
// # worker: 3 : 작업 완료: 작업: 3
// # worker: 3 : 작업 시작: 작업: 8
// # worker: 1 : 작업 완료: 작업: 4
// # worker: 1 : 작업 시작: 작업: 9
// # worker: 2 : 작업 완료: 작업: 6
// # worker: 2 : 작업 시작: 작업: 10
// # worker: 3 : 작업 완료: 작업: 8
// # worker: 3 : 작업 종료!
// # worker: 1 : 작업 완료: 작업: 9
// # worker: 1 : 작업 종료!
// # worker: 4 : 작업 완료: 작업: 7
// # worker: 4 : 작업 종료!
// # worker: 2 : 작업 완료: 작업: 10
// # worker: 2 : 작업 종료!
