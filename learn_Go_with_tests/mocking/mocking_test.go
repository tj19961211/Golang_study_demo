package mocking

import (
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	// 	buffer := &bytes.Buffer{}
	// 	spySleeper := &SpySleeper{}

	// 	Countdown(buffer, spySleeper)

	// 	got := buffer.String()
	// 	want := `3
	// 2
	// 1
	// GO!`

	// 	if got != want {
	// 		t.Errorf("got '%s' want '%s'", got, want)
	// 	}

	// 	if spySleeper.Calls != 4 {
	// 		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	// 	}

	t.Run("Sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
