package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "hello Chris"

		assertMsg(t, got, want)
	})

	t.Run("saying hello world when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "hello world"

		assertMsg(t, got, want)
	})

	t.Run("in japanese", func(t *testing.T) {
		got := Hello("taro", "japanese")
		want := "konnichiwa taro"

		assertMsg(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("hanako", "french")
		want := "bonjour hanako"

		assertMsg(t, got, want)
	})

}