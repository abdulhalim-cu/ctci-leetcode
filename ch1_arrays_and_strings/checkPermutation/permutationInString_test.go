package checkPermutation

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"ab", "eidbaooo", true},
		{"adc", "dcda", true},
		{"abcd", "cbad", true},
		{"abcdef", "dcbefaz", true},
		{"abc", "bca", true},

		// Test cases with non-permutations
		{"ab", "eidboaoo", false},
		{"abc", "def", false},
		{"abc", "ac", false},
		{"abcd", "abc", false},
		{"xyz", "abc", false},

		// Edge cases with equal strings
		{"abc", "abc", true},
		{"abcd", "abcd", true},
		{"abcdef", "abcdef", true},

		// Edge case with s2 being a permutation of s1
		{"abc", "cab", true},
		{"xyz", "zxy", true},

		// Edge case with s1 and s2 having the same characters but different lengths
		{"abcd", "abcdabc", true},
		{"abcd", "abcdabcdabcd", true},
	}

	for _, test := range tests {
		result := CheckInclusion(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("Expected CheckInclusion(%q, %q) to be %v, but got %v", test.s1, test.s2, test.expected, result)
		}
	}
}
