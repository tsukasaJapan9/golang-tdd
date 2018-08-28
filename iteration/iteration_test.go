package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertMsg := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got:%s, want:%s", got, want)
		}
	}

	t.Run("repeat char", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		assertMsg(t, got, want)
	})

	t.Run("repeat char, input count is one", func(t *testing.T) {
		got := Repeat("a", -1)
		want := "a"

		assertMsg(t, got, want)
	})

	t.Run("repeat char, input count is minus", func(t *testing.T) {
		got := Repeat("a", -1)
		want := ""

		assertMsg(t, got, want)
	})

}

func ExampleRepeat() {
	str := Repeat("b", 4)
	fmt.Println(str)
	// output:bbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
