package permutations

import "testing"

func TestFindAnagrams(t *testing.T) {
	// Test Case 1: Basic test case with a single anagram.
	s1 := "cbaebabacd"
	p1 := "abc"
	expected1 := []int{0, 6}
	result1 := findAnagrams(s1, p1)
	if !equalSlices(result1, expected1) {
		t.Errorf("Test Case 1: Expected %v, but got %v", expected1, result1)
	}

	// Test Case 2: Multiple anagrams of "ab" in "abab".
	s2 := "abab"
	p2 := "ab"
	expected2 := []int{0, 1, 2}
	result2 := findAnagrams(s2, p2)
	if !equalSlices(result2, expected2) {
		t.Errorf("Test Case 2: Expected %v, but got %v", expected2, result2)
	}

	// Add more test cases here...
	s3 := "baabbabab"
	p3 := "abb"
	expected3 := []int{2, 3, 4, 6}
	result3 := findAnagrams(s3, p3)
	if !equalSlices(result3, expected3) {
		t.Errorf("Test Case 3: Expected %v, but got %v", expected3, result3)
	}

	// Test Case 10: Anagrams with all lowercase letters.
	s10 := "abcdefghijklmnopqrstuvwxyz"
	p10 := "zxcvbnmasdfghjklqwertyuiop"
	expected10 := []int{0}
	result10 := findAnagrams(s10, p10)
	if !equalSlices(result10, expected10) {
		t.Errorf("Test Case 10: Expected %v, but got %v", expected10, result10)
	}
}

// Helper function to compare slices.
func equalSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
