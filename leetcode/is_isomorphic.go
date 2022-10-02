package leetcode

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	stringMap := make(map[string]string)

	for idx, char := range s {
		thisChar := string(char)
		otherChar := string(t[idx])
		mapValue, exists := stringMap[thisChar]

		// Check if the key is in the map and the value is not the same for both
		if exists && mapValue != otherChar {
			return false
		} else {
			stringMap[thisChar] = otherChar
			stringMap[otherChar] = thisChar
		}
	}

	return true
}
