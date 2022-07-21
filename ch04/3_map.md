## 맵

- 맵은 키:밸류 쌍에 대한 정렬 없는 컬렉션이다.

### 생성과 초기화

```go
dict1 := make(map[string]int)

dict2 := map[string]string{"red": "#da1337", "orange": "#da1332"}
// 맵의 키는 == 연산자를 이용한 비교식에 사용될 수 있는 값이라면 어떤 타입이든 사용할 수 있다.
```

### 맵 활용하기

- 키 밸류 쌍 대입

```go
// 빈 맵 생성
colors := map[string]string{}

colors["red"] = "#da1337"
```

- 초기화를 생략하여 nil 값을 가지는 맵 생성

```go
// 빈 맵 생성
var colors map[string]string

colors["red"] = "#da1337"
//panic: assignment to entry in nil map
```

- 맵에서 값과 키의 존재 여부를 동시에 확인

```go
value, exists := colors["blue"]

if exists {
    fmt.Println(value)
}
// 없다면 value == "", exists == false
```

- 맵에서 값을 조회한 후 확인

```go
value := colors["blue"]

if value != "" {
    fmt.Println(value)
}
// 없다면 value는 제로 값
```

- `for range` 키워드로 맵 반복하기

```go
colors := map[string]string{
    "AliceBlue": "#fsadf",
    "Coral": "#qe1334",
    "DarkGrey": "#erg11",
    "ForestGreen": "#bvf56",
}

for key, value := range colors {
    fmt.Printf("key: %s  value: %s\n", key, value)
}
// key: AliceBlue value: #fsadf
// key: Coral value: #qe1334
// key: DarkGrey value: #erg11
// key: ForestGreen value: #bvf56
```

- 맵에서 아이템 제거하기

```go
// 내장함수 delete
delete(colors, "Coral")
```

- 맵을 함수에 전달하기

```go
// 맵을 전달하면 해당 맵을 참조하는 모든 코드가 영향을 받는다
func removeColor(colors map[string]string, key string) {
    delete(colors, key)
}

func main() {
    colors := map[string]string{
        "AliceBlue": "#fsadf",
        "Coral": "#qe1334",
        "DarkGrey": "#erg11",
        "ForestGreen": "#bvf56",
    }

    removeColor(colors, "Coral")

    for key, value := range colors {
        fmt.Printf("key: %s  value: %s\n", key, value)
    }
}
// key: ForestGreen  value: #bvf56
// key: AliceBlue  value: #fsadf
// key: DarkGrey  value: #erg11
```