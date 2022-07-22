# 메서드

> 2_ex.go 파일 참조

- Go에는 두 종류의 메서드 수신자가 있다.
  - value receiver
  - pointer receiver

```go
func (u user) notify() {
	fmt.Printf("Sending email: %s<%s>\n", u.name, u.email)
}
```

```go
    bill := user{"Bill", "bill@email.com"}
    // value를 통해 value receiver 메서드 호출
    bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
    // pointer를 통해 value receiver 메서드 호출
	lisa.notify()
```

- pointer를 통해 value receiver 메서드를 호출하면 Go는 포인터 값을 역참조하여 value receiver에 정의된 메서드를 호출한다.

```go
    (*lisa).notify()
```

- 반대로 value를 통해 point receiver 메서드를 호출하면, 값을 참조하여 메서드 호출에 적합한 receiver 타입으로 변환한다.

```go
    (&bill).changeEmail("bill@newmail.com")
```

