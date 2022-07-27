## 벤치마크 실행

- `-v` : 상세한 결과 출력 
- `-run="none"` : 단위 테스트 실행 스킵
- `-bench` : 벤치마크 실행 대상 함수. `.` 은 모두 실행
- `-benchtime` : 벤치마킹 시간 설정 (없으면 기본 1초간 실행). 너무 크게 잡으면 무의미해진다. 
- `-benchmem` : 힙메모리 할당 횟수와 한 번 할당할 때의 크기를 제공

```shell
go test -v -run="none" -bench="BenchmarkSprintf"

go test -v -run="none" -bench="BenchmarkSprintf" -benchtime=3s

go test -v -run="none" -bench=. -benchtime=3s

go test -v -run="none" -bench=. -benchtime=3s -benchmem
```