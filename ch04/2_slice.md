## 슬라이스

### 슬라이스 내부 구조

- Go가 내부의 배열을 조작하는데 필요한 메타데이터를 관리하는 세 개의 필드로 구성된다.

```
데이터: 메모리주소   3     5
메타:   포인터     길이   용량

[0] [1] [2] [3] [4]
10  20  30   0   0
```

### 생성과 초기화

```go
// 길이를 명시
// 길이=5, 용량=5
slice1 := make([]string, 5)

// 길이와 용량을 명시
// 길이=3, 용량=5
slice2 := make([]int ,3, 5)

// 길이보다 작은 크기의 용량을 지정하면 컴파일 에러
slice3 := make([]int, 5, 3)
// invalid argument: length and capacity swapped
```

```go
// 슬라이스 리터럴을 이용한 생성
// 길이=5, 용량=5
slice1 := []string{"red", "blue", "green", "yello", "pink"}
// 인덱스로 슬라이스 생성
slice2 := []string{99: ""}
```

- `[]` 연산자에 값을 지정하면 배열, 지정하지 않으면 슬라이스가 생성된다.

```go
array := [3]int{10, 20, 30}
slice := []int{10, 20 ,30}
```

- nil 슬라이스

```go
var slice []int
//데이터:  nil    0     0
//메타:   포인터 길이   용량
```

- 빈 슬라이스

```go
// make 함수로 빈 정수 슬라이스
slice := make([]int, 0)
// 슬라이스 표현식으로 빈 정수 슬라이스
slice := []int{}
//데이터: 메모리주소    0     0
//메타:   포인터      길이   용량
```

### 슬라이스 활용

- 배열 표현식을 이용해 값 대입하기

```go
// 길이=5, 용량=5
slice := []int{10, 20, 30, 40, 50}
slice[1] = 25
// 길이=2, 용량=4
newSlice := slice[1:3]
// slice:    | addr0/5/5
//           10 25 30 40 50
// newSlice:    | addr1/2/4

// newSlice는 메모리를 slice와 공유한다
newSlice[1] = 35
// slice[2] == 35
```

- 용량이 k인 내부 배열을 갖는 슬라이스에 대한 slice[i:j] 연산의 결과
  - 길이 = j - i
  - 용량 = k - i

- 인덱스가 범위를 벗어나면 런타임 에러

```go
slice := []int{10, 20, 30, 40, 50}
newSlice := slice[1:3]
newSlice[3] = 45
//panic: runtime error: index out of range [3] with length 2
```

- 슬라이스 크기 확장

```go
slice := []int{10, 20, 30, 40, 50}
newSlice := slice[1:3]
// append 함수는 변경된 새로운 슬라이스를 리턴한다.
// newSlice는 아직 여분의 용량이 남아있기 때문에 append할 수 있다.
newSlice = append(newSlice, 60)
// slice:    | addr0/5/5
//           10 20 30 60 50
// newSlice:    | addr1/3/4
```

- 슬라이스의 용량이 부족하다면 append 함수는 새로운 내부 배열을 생성하고 기존 값들을 새로운 배열로 복사한 후 값을 추가한다.
  - 기존 슬라이스 용량이 1,000 보다 작은 경우 두 배로 확장되고 1,000개를 넘으면 25%씩 증가한다.

```go
slice := []int{10, 20, 30, 40}
//slice: 길이=4, 용량=4
//10 20 30 40

newSlice = append(newSlice, 50)
//newSlice: 길이=5, 용량=5
//10 20 30 40 50 0 0 0


newSlice2 = append(newSlice, 60)
//newSlice2: 길이=6, 용량=6
//10 20 30 40 50 60 0 0
```

- 세번째 인덱스를 이용해 슬라이스 자르기

```go
source := []string{"apple", "orange", "plum", "banana", "grape"}
slice := source[2:3:4]
// len=1, cap=2
```

- slice[i:j:k]
  - 길이 = j - i
  - 용량 = k - i

- 길이와 용량을 동일하게 설정하는 장점

```go
source := []string{"apple", "orange", "plum", "banana", "grape"}
slice := source[2:3:3]
// len=1, cap=1
slice = append(slice, "kiwi")
// 새 슬라이스의 용량을 지정하지 않으면 슬라이스가 용량을 조절할 떄 내부 배열의 남은 공간을 계속 참조하므로, "banana"가 지워지고 kiwi가 들어간다.
// 용량을 1로 제한하기 때문에, kiwi를 append할때 새로운 내부배열을 참조하는 새로운 슬라이스를 생성한다.
// plum kiwi
```

- append는 가변함수여서 ... 연산자를 쓸 수 있다.

```go
s1 := []int{1, 2}
s2 := []int{3, 4}

fmt.Printf("%v\n", append(s1, s2...))
[1 2 3 4]
```

- `for range` 키워드로 슬라이스 반복하기

```go
slice := []int{10, 20, 30, 40}

for index, value := range slice {
  fmt.Printf("index: %d value: %d\n", index, value)
}
```

- `range` 키워드는 값에 대한 참조를 리턴하는 것이 아니라 복사본을 생성한다.(인덱스&배열원소값) 따라서 각 원소값을 저장한 변수의 주소를 포인터로 사용하면 실수가 생긴다.

```go
slice := []int{10, 20, 30, 40}

for index, value := range slice {
  fmt.Printf("value: %d addrOfValue: %X addrOfElem: %X\n", value, &value, &slice[index])
}

// value: 10 addrOfValue: C0000BA000 addrOfElem: C0000B8000
// value: 20 addrOfValue: C0000BA000 addrOfElem: C0000B8008
// value: 30 addrOfValue: C0000BA000 addrOfElem: C0000B8010
// value: 40 addrOfValue: C0000BA000 addrOfElem: C0000B8018

// value변수가 값의 복사본을 저장하는 변수이기 때문에 addrOfValue는 항상 동일한 값을 가진다.
// index 변수도 동일
```

- `range` 키워드는 무조건 첫번째 요소부터 루프를 돈다. 세밀한 조정은 `for` 루프를 사용하면 된다.

```go
slice := []int{10, 20, 30, 40}

for index := 2; index < len(slice); index++ {
  fmt.Printf("index: %d value: %d\n", index, slice[index])
}
```

### 슬라이스를 함수에 전달하기

- 함수에 슬라이스를 전달하려면 그냥 슬라이스를 전달하기만 하면 된다. 슬라이스의 크기는 매우 작아서 이를 복사하여 함수 간에 전달하는 처리 비용이 크지 않다.
  - 포인터 8바이트 + 길이 8바이트 + 용량 8바이트 = 24바이트
  - 슬라이스만 복사되고 내부 배열은 복사되지 않는다.

```go
slice := make([]int, 1e6)

slice = foo(slice)

func foo(slice []int) []int {
  // ...
  return slice
}
```