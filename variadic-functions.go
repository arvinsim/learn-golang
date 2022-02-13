package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//func printOutMessages(messages ...string) {
//	for _, message := range messages {
//		fmt.Println(message)
//	}
//}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//printOutMessages([]string{"Hey there", "Hallo!", "Hi"}...)
}
