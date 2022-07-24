package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// sharing resource
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go increaseCounter(1)
	go increaseCounter(2)

	wg.Wait()
	fmt.Println("result: ", counter)
}

func increaseCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Add 1 to the counter
		atomic.AddInt64(&counter, 1)
		// Gosched yields the processor, allowing other goroutines to run.
		// It does not suspend the current goroutine, so execution resumes automatically.
		runtime.Gosched()
	}
}

// result:  4
// atomic function을 쓰지 않았다면 경쟁상태에 빠지고 결과는 2가 되었을 것이다.
