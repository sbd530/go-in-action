# 로깅

GO의 표준 라이브러리는 약간의 설정만으로 편리하게 활용할 수 있는 `log` 패키지를 제공한다. 물론 커스텀 로거를 작성하여 필요한 로깅 기능을 구현할 수도 있다.

## log 패키지

```go
package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
    // 로그 설정을 비트 OR을 통해 조정할 수 있다.
    // 2022/07/26 01:12:34.123123 /a/b/c/d.go:23: 메시지
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 은 표준 로거에 메시지를 출력한다.
	log.Println("Message")
	// Fatalln 은 Println() 을 실행한 후 os.Exit(1)을 호출한다.
	log.Fatalln("Fatal err msg")
	// Panicln 은 Println() 을 호출한 후 panic() 을 호출한다.
	log.Panicln("Panic msg")
}
```

- `init()` 함수 안에서 로그 관련 설정을 구성한다.

- `log.go` 파일을 들어가 살펴보면 다음과 같이 플래그값에 대해 알 수 있다.
  - `iota` 키워드는 대입 구문이 나올때까지 동일한 표현식을 매 상수마다 중복적용할 것을 컴파일러에게 알린다.
  - `iota` 키워드는 디폴트값 0부터 시작하여 상수 정의마다 1씩 증가시키는 것에도 사용된다.

```go
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
// 비트 연산
// 1 << 0 = 1
// 1 << 1 = 2
// 1 << 2 = 4
// 1 << 3 = 8
// 1 << 4 = 16
```

## 사용자정의 로거

```go
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("에러 로그 파일을 열 수 없다.", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("일반적인 메시지")
	Info.Println("특별한 정보")
	Warning.Println("경고 메시지")
	Error.Println("에러 메시지")
}
```

- `log.New()` 함수 파라미터로 `io.Writer`를 받고 있어서 다양한 Writer 구현체가 들어갈 수 있다.

```go
// golang.org/src/log/log.go
// New함수는 새로운 로거를 생성한다.
// out 변수는  로그 데이터가 기록될 곳을 지정하기위한 파라미터
// prefix 는 로그생성시 각죽 시작에 삽입될 문자열
// flag는 로그의 속성을 정의하는 파라미터
func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out: out, prefix: prefix, flag: flag}
}
```