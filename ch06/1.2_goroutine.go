package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	// 스케줄러가 사용할 하나의 논리 프로세서를 할당
	runtime.GOMAXPROCS(1)

	// wg는 프로그램의 종료를 대기하기 위해 사용

	// 총 두개의 카운트를 추가
	wg.Add(2)

	fmt.Println("# 고루틴 실행")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("# 대기중...")
	wg.Wait()

	fmt.Println("\n# 프로그램 종료")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("완료: ", prefix)
}
