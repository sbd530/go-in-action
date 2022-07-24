package main

import (
	"fmt"

	"./entities"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	// 비노출 타입인 내부 타입의 노출타입 필드에 값을 대입
	a.Name = "Bill"
	a.Email = "bill@email.com"
	// 내부 타입의 식별자들은 외부 타입으로 승격되기 때문에 접근할 수 있다.

	fmt.Printf("User: %v\n", a)
}
