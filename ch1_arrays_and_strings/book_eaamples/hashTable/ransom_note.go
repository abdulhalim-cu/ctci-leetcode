package hashTable

func canConstruct(ransomNote string, magazine string) bool {
	freq := [26]int{}

	for i := 0; i < len(magazine); i++ {
		freq[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		freq[ransomNote[i]-'a']--
	}

	for _, v := range freq {
		if v < 0 {
			return false
		}
	}
	return true
}
