package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

// value receiver
func (u user) notify() {
	fmt.Printf("Sending email: %s<%s>\n", u.name, u.email)
}

// pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	// user 타입의 value를 통해 value receiver 메서드 호출
	bill := user{"Bill", "bill@email.com"}
	bill.notify()
	// user 타입의 pointer를 통해 value receiver 메서드 호출
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 타입의 value를 통해 pointer receiver 메서드 호출
	bill.changeEmail("bill@newmail.com")
	bill.notify()
	// user 타입의 pointer를 통해 pointer receiver 메서드 호출
	lisa.changeEmail("lisa@newmail.com")
	lisa.notify()
}

// Sending email: Bill<bill@email.com>
// Sending email: Lisa<lisa@email.com>
// Sending email: Bill<bill@newmail.com>
// Sending email: Lisa<lisa@newmail.com>
