package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// 스케줄러가 사용할 하나의 논리 프로세서를 할당
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2) // 병렬적으로 실행

	// wg는 프로그램의 종료를 대기하기 위해 사용
	var wg sync.WaitGroup
	// 총 두개의 카운트를 추가
	wg.Add(2)

	fmt.Println("# 고루틴 실행")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("# 대기중...")
	wg.Wait()

	fmt.Println("\n# 프로그램 종료")
}

// # 고루틴 실행
// # 대기중...
// A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z a b c d e f g h i j k l m n o p q r s t u v w x y z
// # 프로그램 종료

// 연산시간이 너무 짧아서 스케줄을 할당받고 대문자를 모두 출력한 후 다음 고루틴이 실행되었따.
