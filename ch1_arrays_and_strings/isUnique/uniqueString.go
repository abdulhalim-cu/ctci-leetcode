package isunique

func CheckHasUniqueChar(s string) bool {
	if len(s) > 256 {
		return false
	}
	charSet := make(map[rune]bool, 256)
	for _, c := range s {
		if charSet[c] {
			return false
		}
		charSet[c] = true
	}
	return true
}

// assume that all inputs are lower case english letter from a-z; in that case we can do perform bitwise operation to
// check the duplication of letters

func CheckHasUniqueCharUsingBitwiseOperation(s string) bool {
	checker := 0
	for i := 0; i < len(s); i++ {
		val := s[i] - 'a'
		if checker&(1<<val) > 0 {
			return false
		}
		checker |= 1 << val
	}
	return true
}
