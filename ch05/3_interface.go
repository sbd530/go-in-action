package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

type admin struct {
	name  string
	email string
}

// pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending mail: %s<%s>\n", u.name, u.email)
}

func (u *admin) notify() {
	fmt.Printf("Sending mail: %s<%s>\n", u.name, u.email)
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	// 인테페이스 타입 인수에 value 전달
	// sendNotification(u)
	// ./prog.go:23:19: cannot use u (variable of type user) as type notifier in argument to sendNotification:
	// user does not implement notifier (notify method has pointer receiver)

	// 인테페이스 타입 인수에 pointer 전달
	sendNotification(&bill)

	// 인터페이스를 이용한 다형성
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
