package main

import (
	"testing"
	"bytes"
	"reflect"
	"time"
)

const sleepDesc = "sleep"
const writeDesc = "write"

type CountdownOperationsSpy struct {
	calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.calls = append(s.calls, sleepDesc)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.calls = append(s.calls, writeDesc)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	assertDuration := func(t *testing.T, got, want time.Duration) {
		t.Helper()
		if got != want {
			t.Errorf("got:%v, want:%v", got, want)
		}
	}

	t.Run("output string test", func(t *testing.T) {
		buf := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}
		Countdown(buf, spySleeper)

		got := buf.String()
		want := "3\n2\n1\nGo!\n"

		assert(t, got, want)
	})

	t.Run("sleep order test", func(t *testing.T) {
		spySleeper := &CountdownOperationsSpy{}

		Countdown(spySleeper, spySleeper)

		got := spySleeper.calls
		want := []string{
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got:%v, want:%v", got, want)
		}
	})

	t.Run("sleep duration test", func(t *testing.T) {
		sleepTime := 5 * time.Second

		spyTime := &SpyTime{}
		spySleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}

		spySleeper.Sleep()

		assertDuration(t, spyTime.durationSlept, sleepTime)
	})
}