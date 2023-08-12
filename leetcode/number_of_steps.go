package leetcode

func NumberOfSteps(n int) int {
	i := 0
	for n > 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n - 1
		}
		i++
	}

	return i
}
