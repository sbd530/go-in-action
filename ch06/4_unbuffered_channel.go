// 2개의 고루틴을 사용한 테니스경기 메타포
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 버퍼가 없는 정수 채널 생성 (테니스 코트)
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	// 경기 시작
	court <- 1

	// 경기 종료 대기
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		// 공이 돌아올 때까지 대기
		ball, ok := <-court
		if !ok {
			// 채널이 닫혔으면 승리한 것으로 간주
			fmt.Printf("%s 선수가 승리!\n", name)
			return
		}

		// 랜덤 양의 정수을 이용해 공을 받아치지 못했는지 확인
		n := rand.Intn(100)
		fmt.Printf("# n=%d\n", n)
		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 못받아침...\n", name)

			// 채널을 닫아서 현재 선수가 패배했다고 알린다.
			close(court)
			return
		}

		// 선수가 공을 받아친 횟수를 찍고 증가시킨다.
		fmt.Printf("%s 선수가 %d 번째 공을 받아쳤음~\n", name, ball)
		ball++

		// 상대에게 공을 보낸다.
		court <- ball
	}
}
