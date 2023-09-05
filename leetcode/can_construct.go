package leetcode

import (
	"strings"
)

func CanConstruct(ransomNote string, magazine string) bool {
	for _, char := range magazine {
		ransomNote = strings.Replace(ransomNote, string(char), "", 1)
	}
	return len(ransomNote) == 0
}
