package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	// 임계 지역을 설정할 떄 사용할 뮤텍스
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("# result: %d\n", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 락을 획득하여 한번에 하나의 고루틴만 이 임계지역에 접근하도록 한다.
		mutex.Lock()
		// 임계지역을 쉽게 볼수있도록 만든 {}. 반드시 블록이 필요한 건 아니다.
		{
			// counter 값을 읽는다.
			value := counter
			// 스레드를 양보해서 큐로 돌아가도록 한다.
			// 양보하더라도 스케줄러는 동일한 고루틴을 할당해 계속해서 실행되도록 한다.
			runtime.Gosched()
			// 값을 증가
			value++
			// 원래 counter에 값을 저장
			counter = value
		}
		// 잠금해제 -> 다른 고루틴이 접근
		mutex.Unlock()
	}
}

// # result: 4
