package permutations

func uniqPermutations(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}

	visited := make(map[string]bool)
	return uniqPermutation(s, visited)
}

func uniqPermutation(s string, visited map[string]bool) []string {
	permutations := []string{}

	for i := 0; i < len(s); i++ {
		sToPermute := s[:i] + s[i+1:]
		if !visited[sToPermute] {
			restPermutations := uniqPermutations(sToPermute)
			for _, perm := range restPermutations {
				permutations = append(permutations, string(s[i])+perm)
			}
		}
		visited[sToPermute] = true
	}
	return permutations
}
