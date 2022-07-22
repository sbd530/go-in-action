# 사용자 정의 타입

```go
// 시스템에 user 타입을 선언
type user struct {
    name       string
    email      string
    ext        int
    privileged bool
}

// user 타입의 변수를 선언. 제로 값으로 초기화
var bill user

// 구조체 리터럴을 통해 구조체 타입의 변수를 선언
lisa := user {
    name:       "Lisa",
    email:      "lisa@email.com",
    ext:        123,
    privileged: true,
}
// 필드이름 없이
lisa2 := user{"Lisa","lisa@email.com",123,true}
```

```go
// int64 을 base type으로 타입 선언
type Duration int64

// base type이 int64이지만 둘은 별개의 타입이기 때문에 할당하려면 에러 발생
var dur Duration
dur = int64(1000)
// ./prog.go:9:8: cannot use int64(1000) (constant 1000 of type int64) as type Duration in assignment
```