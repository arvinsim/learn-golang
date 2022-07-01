package main

import (
	"fmt"
	"learn-golang/leetcode"
)

func main() {
	// Should be false
	result := leetcode.IsIsomorphic("badc", "bada")

	fmt.Printf("%t", result)
}
