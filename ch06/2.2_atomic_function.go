package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// 실행 중인 고루틴들의 종료신호로 사용되는 플래그
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")
	// 고루틴이 실행될 시간을 확보한다.
	time.Sleep(1 * time.Second)

	fmt.Println("# 프로그램 종료")
	// 원자적으로 플래그를 종료 상태로 저장
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("# 작업 진행중: %s\n", name)
		time.Sleep(250 * time.Millisecond)
		// 종료 플래그를 확인하고 루프를 빠져나온다.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("# 작업 종료: %s\n", name)
			break
		}
	}
}

// # 작업 진행중: A
// # 작업 진행중: B
// # 작업 진행중: B
// # 작업 진행중: A
// # 작업 진행중: B
// # 작업 진행중: A
// # 작업 진행중: A
// # 작업 진행중: B
// # 작업 진행중: A
// # 프로그램 종료
// # 작업 종료: B
// # 작업 종료: A
