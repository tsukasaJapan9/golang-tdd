package arrays

// Sum is return sum of input ary
func Sum(ary []int) int {
	sum := 0

	for _, n := range ary {
		sum += n
	}

	return sum
}

// SumAll  is return sum of input multi ary
func SumAll(numbersToSum ...[]int) ([]int) {
	var sums []int

	for _, n := range numbersToSum {
		sums = append(sums, Sum(n))
	}

	return sums
}

// SumAllTails is return sum of input ary tails
func SumAllTails(numbersToSum ...[]int) ([]int) {
	var sums []int

	for _, n := range numbersToSum {
		if len(n) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(n[1:]))
		}
	}

	return sums
}