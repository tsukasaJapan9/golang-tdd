package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

const countStart = 3
const finalWord = "Go"

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(d time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(writer, "%d\n", i)
	}

	sleeper.Sleep()
	fmt.Fprintf(writer, "%s!\n", finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{duration:1*time.Second, sleep: time.Sleep}

	Countdown(os.Stdout, sleeper)
}