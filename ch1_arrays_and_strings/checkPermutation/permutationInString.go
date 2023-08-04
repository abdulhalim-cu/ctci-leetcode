package checkPermutation

func CheckInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Count, s2Count := make([]int, 26), make([]int, 26)
	for i := 0; i < len(s1); i++ {
		s1Count[s1[i]-'a']++
		s2Count[s2[i]-'a']++
	}

	matchedCount := 0
	for i := 0; i < 26; i++ {
		if s1Count[i] == s2Count[i] {
			matchedCount++
		}
	}

	l := 0
	for r := len(s1); r < len(s2); r++ {
		if matchedCount == 26 {
			return true
		}
		idx := s2[r] - 'a'
		s2Count[idx]++
		if s1Count[idx] == s2Count[idx] {
			matchedCount++
		} else if s1Count[idx]+1 == s2Count[idx] {
			matchedCount--
		}
		idx = s2[l] - 'a'
		s2Count[idx]--
		if s1Count[idx] == s2Count[idx] {
			matchedCount++
		} else if s1Count[idx]-1 == s2Count[idx] {
			matchedCount--
		}
		l++
	}
	return matchedCount == 26
}
