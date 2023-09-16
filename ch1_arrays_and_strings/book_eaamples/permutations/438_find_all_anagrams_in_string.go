package permutations

//func findAnagrams(s string, p string) []int {
//	var res = []int{}
//	lnS, lnP := len(s), len(p)
//	if lnP > lnS {
//		return res
//	}
//	for i := 0; i <= (lnS - lnP); i++ {
//		if isAnagram(s[i:i+lnP], p) {
//			res = append(res, i)
//		}
//	}
//	return res
//}
//
//func isAnagram(s, p string) bool {
//	if len(s) != len(p) {
//		return false
//	}
//	var refs = make([]int, 26)
//	for i := 0; i < len(s); i++ {
//		refs[s[i]-'a']++
//		refs[p[i]-'a']--
//	}
//	for _, v := range refs {
//		if v != 0 {
//			return false
//		}
//	}
//	return true
//}

func findAnagrams(s string, p string) []int {
	var res = []int{}
	lnS, lnP := len(s), len(p)
	if lnS < lnP {
		return res
	}
	freqS := [26]int{}
	freqP := [26]int{}
	start, end := 0, 0

	for i := 0; i < lnP; i++ {
		freqP[p[i]-'a']++
	}

	for end < lnS {
		freqS[s[end]-'a']++
		if end-start+1 == lnP {
			if freqS == freqP {
				res = append(res, start)
			}
			freqS[s[start]-'a']--
			start++
		}
		end++
	}
	return res
}
