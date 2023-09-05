package practice

func SumOfIntegers(nums ...int) int {
	var accumulator = 0

	for _, num := range nums {
		accumulator += num
	}

	return accumulator
}

func IsOddOrEven(inputNumber int) string {
	var isEven = inputNumber%2 == 0

	if isEven {
		return "even"
	}

	return "odd"
}
