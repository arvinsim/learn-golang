package leetcode

func GetRichestCustomerWealth(customerMatrix [][]int) int {
	biggestWealth := 0

	for _, customer := range customerMatrix {
		customerWealth := 0
		for _, wealth := range customer {
			customerWealth += wealth
		}
		if customerWealth > biggestWealth {
			biggestWealth = customerWealth
		}
	}

	return biggestWealth
}
