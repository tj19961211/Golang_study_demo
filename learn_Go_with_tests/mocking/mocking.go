package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

/*虽然我们都知道 *bytes.Buffer 可以运行，但最好使用通用接口代替。*/
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

// type ConfigurableSleeper struct {
// 	duration time.Duration
// }

// func (o *ConfigurableSleeper) Sleep() {
// 	time.Sleep(o.duration)
// }

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"
