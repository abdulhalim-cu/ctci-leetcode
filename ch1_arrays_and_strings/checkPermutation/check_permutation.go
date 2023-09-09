package checkPermutation

// sorting and comparing is not an optimal solution

/*
func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	t1 := []rune(s1)
	t2 := []rune(s2)
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})

	sort.Slice(t2, func(i, j int) bool {
		return t2[i] < t2[j]
	})
	return string(t1) == string(t2)
}
*/

// optimal solution is counting character

func IsPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	charCount := make(map[rune]int)

	for _, c := range s1 {
		charCount[c]++
	}
	for _, c := range s2 {
		charCount[c]--
		if charCount[c] < 0 {
			return false
		}
	}
	return true
}
