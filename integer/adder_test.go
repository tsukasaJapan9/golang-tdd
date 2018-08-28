package integer

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	assertMsg := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got:%d, want:%d", got, want)
		}
	}

	t.Run("add 2 integer", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertMsg(t, got, want)
	})

	t.Run("add 2 integer part2", func(t *testing.T) {
		got := Add(3, 5)
		want := 8

		assertMsg(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// output:6
}