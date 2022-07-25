package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// 주어진 타임아웃 시간 동안 작업을 수행한다.
// OS 인터럽트에 의해 실행이 종료된다.
type Runner struct {
	// 인터럽트 신호를 수신하기 위한 채널
	interrupt chan os.Signal

	// 처리가 종료되었을을 알리기 위한 채널
	complete chan error

	// 지정된 시간이 초과했음을 알리기 위한 채널
	timeout <-chan time.Time

	// 인덱스 순서로 처리될 작업의 목록을 저장하기 위한 슬라이스
	tasks []func(int)
}

// timeout 채널에서 값 수신시 ErrTimeout 리턴
var ErrTimeout = errors.New("시간 초과!")

// OS 이벤트 수신시 ErrInterrupt 리턴
var ErrInterrupt = errors.New("OS 인터럽트 신호를 수신했습니다.")

// 실행할 Runner 타입 값을 리턴하는 함수
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Runner 타입에 작업을 추가하는 메서드
// 작업은 int형 ID를 매개변수로 전달받는 함수이다.
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// 저장된 모든 작업을 실행하고 채널 이벤트를 관찰
func (r *Runner) Start() error {
	// 모든 종류의 인터럽트 수신
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// 작업 완료 신호 수신
	case err := <-r.complete:
		return err
	// 작업시간 초과 수신
	case <-r.timeout:
		return ErrTimeout
	}
}

// 개별 작업 실행
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 작업 실행
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		// 이후 발생하는 인터럽트 신호를 더이상 수신하지 않는다.
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
