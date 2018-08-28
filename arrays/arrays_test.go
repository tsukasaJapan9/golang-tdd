package arrays

import (
	"testing"
	"reflect"
)

func TestSum(t *testing.T) {
	assetMsg := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got:%d, want:%d, given: %v", got, want, numbers)
		}
	}
	
	t.Run("sum 5 array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got  := Sum(numbers)
		want := 15

		assetMsg(t, got, want, numbers)
	})

	t.Run("sum 4 array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got  := Sum(numbers)
		want := 10

		assetMsg(t, got, want, numbers)
	})

}

func TestSumAll(t *testing.T) {
	assetMsg := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got:%v, want:%v", got, want)
		}
	}
	t.Run("sum 2 array", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{3, 4, 5}

		got  := SumAll(numbers1, numbers2)
		want := []int{6, 12}

		assetMsg(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	assetMsg := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got:%v, want:%v", got, want)
		}
	}
	t.Run("sum 2 array of tail", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{3, 4, 5}

		got  := SumAllTails(numbers1, numbers2)
		want := []int{5, 9}

		assetMsg(t, got, want)
	})

	t.Run("sum 2 array of tail", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{}

		got  := SumAllTails(numbers1, numbers2)
		want := []int{5, 0}

		assetMsg(t, got, want)
	})

}