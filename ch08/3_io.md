# 입력과 출력

UNIX 기반 운영체제의 장점 중 하나는 한 프로그램의 출력을 다른 프로그램의 입력으로 사용할 수 있다는 것이다. 이 철학 덕분에 한 가지 일에만 집중하며 그것만 잘 처리하는 작고 간단한 프로그램들의 집합이 구성될 수 있었다. 도한 이로 인해 이 프로그램들을 조합해서 더 놀라운 일들을 수행하는 스크립트를 작성할 수 있었다. 이 과정에서 `stdout`과 `stdin` 장치는 프로세스 사이의 데이터를 교환하기 위한 통로로 동작한다.

이와 동일한 개념이 `io` 패키지에도 그대로 옮겨져있으며, 이 패키지가 제공하는 기능은 실로 놀라운 수준이다. 이 패키지는 데이터의 타입, 데이터의 출처 또는 그 목적지와는 무관하게 데이터의 스트림을 매우 효과적으로 처리할 수 있다. 이를 위해 `stdout`과 `stdin` eotls `io.Writer`와 `io.Reader` 인터페이스를 사용한다. 이 인터페이스를 구현하는 타입들을 이용하면 `io` 패키지가 제공하는 기능은 물론, 다른 패키지에 선언된 함수와 메서드 중 이 인터페이스를 매개변수로 사용하는 모든 것을 활용할 수 있다. 바로 이점이 인터페이스 타입을 바탕으로 기능과 API를 구현했을 때 얻을 수 있는 장점이다.

## Writer 인터페이스와 Reader 인터페이스

- `io.Writer` 인터페이스를 선언한 코드

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
// p에서 len(p) 만큼의 데이터를 스트림에 출력한다.
// p에서 출력된 바이트의 길이(0<=n<=len(p))와 쓰기 동작의 중단을 유발하는 에러가 있는 경우 에러를 함께 리턴한다.
// Write 메서드는 출력된 바이트의 길이 n이 len(p)보다 작을 경우 반드시 nil이 아닌 에러를 리턴해야 한다.
// 어떤 경우라도 슬라이스 내의 데이터를 변경해서는 안된다. 
```

- `io.Reader` 인터페이스를 선언한 코드

```go
type Reader interface {
    Reade(p []byte) (n int, err error)
}
// p로부터 len(p) 만큼의 데이터를 읽는다.
// 읽어들인 바이트의 길이(0<=n<=len(p))를 리턴하며, 오류가 발생한 경우 오류를 함께 리턴한다.
// 길이 n이 len(p)보다 작을 경우 p의 데이터를 모두 읽었지만 처리 과정에서 p의 공간이 늘어난 경우일 수 있다.
// 만일 p에 사용 가능한 데이터가 있지만 그 길이가 len(p)보다 작다면 Read메서드는 규칙에 따라 p에 데이터가 채워지기를 기다리지 않고 현재 사용가능한 값만을 리턴한다.
```

- 읽기 도중 파일의 끝에 도달한 경우, 마지막 바이트를 읽었을 때 선택 가능한 옵션은 두가지다. 마지막 바이트를 읽은후 EOF를 에러값으로 리턴하거나 마지막 바이트를 읽은 후 nil 값을 에러 값으로 리턴하는 것이다. 후자라면 다음번 Read 메서드를 호출할 때 읽은 바이트 길이는 0을, 에러값은 EOF를 리턴해야한다.