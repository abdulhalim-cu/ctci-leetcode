package checkPermutation

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var refs = make([]int, 256)
	for i := 0; i < len(s); i++ {
		refs[s[i]]++
		refs[t[i]]--
	}
	for _, v := range refs {
		if v != 0 {
			return false
		}
	}
	return true
}
