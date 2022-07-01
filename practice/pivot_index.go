package practice

func main() {
}

func PivotIndex(nums []int) int {
	for idx, _ := range nums {
		leftSum := 0
		rightSum := 0

		for _, element := range nums[:idx] {
			leftSum += element
		}

		for _, element := range nums[idx+1:] {
			rightSum += element
		}

		if leftSum == rightSum {
			return idx
		}
	}

	return -1
}
