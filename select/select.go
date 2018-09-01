package main

import (
	"time"
	"net/http"
	"fmt"
)

func measureTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

var tenTimeout = 10 * time.Millisecond

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timeout")
	}
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenTimeout)
}

func main() {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	got, _ := Racer(slowURL, fastURL)
	fmt.Println("faster:" + got)
}
