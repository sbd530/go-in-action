package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending mail: %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // 포함된 타입
	level string
}

func main() {
	adm := admin{
		user: user{
			name:  "jonh",
			email: "john@email.com",
		},
		level: "super",
	}

	// call inner type's method
	adm.user.notify()
	// inner type's method promotion
	adm.notify()
}
