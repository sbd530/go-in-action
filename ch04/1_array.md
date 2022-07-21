## 배열

### 배열 선언과 초기화

```go
// 배열을 제로 값으로 초기화하여 선언
var array [5]int

// 배열 리터럴 이용
array := [5]int{10, 20, 30, 40, 50}

// 자동으로 길이가 결정되는 배열 선언
array := [...]int{10, 20, 30, 40, 50}

// 일부 요소만 초기화하는 배열 선언
array := [5]int{1: 10, 2: 20}
```

### 배열 활용

- 배열 원소에 접근하기

```go
array := [5]int{10, 20, 30, 40, 50}
// 원소 값 변경
array[2] = 35
```

- 원소의 포인터 요소에 접근하기

```go
// 정수 포인터 배열 선언, 인덱스 0,1을 정수 포인터로 초기화
// 초기화 안된 원소는 nil
array := [5]*int{0: new(int), 1: new(int)}

// 값 대입
*array[0] = 10
*array[1] = 20
// addr0 addr1 nil nil nil
//  10    20 
```

- 배열을 같은 타입의 다른 배열에 대입하기
  - Go에서 배열은 값으로 취근된다.

```go
var array1 [5]string
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

// array2의 값을 array1로 복사한다.
array1 = array2
// array2 : "Red", "Blue", "Green", "Yellow", "Pink"
// array1 : "Red", "Blue", "Green", "Yellow", "Pink" -> 서로 다른 메모리 공간에 각각 저장
```

- 포인터 배열을 다른 배열에 복사하기

```go
var array1 [3]*string
array2 := [3]*string{new(string), new(string), new(string)}

*array2[0] = "Red"
*array2[1] = "Blue"
*array2[2] = "Green"

array1 = array2

// array2 : addr0 addr1 addr2
//          Red   Blue  Green   
// array1 : addr0 addr1 addr2
```

### 다차원 배열

- 이차원 배열 선언

```go
var array1 [4][2]int
array2 := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
array3 := [4][2]int{1: {20, 21}, 3: {40, 41}}
array4 := [4][2]int{1: {0: 20}, 3: {1: 41}}
```

- 이차원 배열의 원소에 접근하기

```go
var array [2][2]int

array[0][0] = 10
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40
```

- 동일 타입의 다차원 배열 대입하기

```go
var array1 [2][2]int
var array2 [2][2]int

array2[0][0] = 10
array2[0][1] = 20
array2[1][0] = 30
array2[1][1] = 40

array1 = array2

// 배열은 값으로 취급되기 때문에 각 차원을 개별적으로 복사할 수도 있다.
var array3 [2]int = array[1]

var value int = array1[1][0]
```

### 함수에 배열 전달하기

- 함수간에 변수를 전달할 때는 항상 그 값이 전달되기 때문에 함수에 배열을 전달하는 것은 크기와 상관없이 배열 전체를 복사하여 함수에 전달하기 때문에 많은 비용이 소모된다.

```go
var array [1e6]int

foo(array)
// 백만개의 정수 배열을 매개변수로 전달
// 함수 호출시마다 8메가바이트의 메모리가 스택에 할당된다.
func foo(array [1e6]int) { ... }
```

- 함수에 배열을 포인터로 전달하면 스택에는 포인터 변수를 위한 8바이트 크기의 메모리만 할당된다.

```go
var array [1e6]int

foo(&array)

func foo(array *[1e6]int) { ... }
```