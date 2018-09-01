package depinj

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	t.Run("first test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "hello Chris"

		assert(t, got, want)
	})
}