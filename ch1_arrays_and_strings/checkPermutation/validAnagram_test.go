package checkPermutation

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		// Test cases with anagrams
		{"anagram", "nagaram", true},
		{"listen", "silent", true},
		{"test", "sett", true},

		// Test cases with non-anagrams
		{"hello", "world", false},
		{"abc", "abcd", false},
		{"abcdef", "efghij", false},

		// Edge case with empty strings
		{"", "", true},
	}

	for _, test := range tests {
		result := isAnagram(test.s, test.t)
		if result != test.expected {
			t.Errorf("Expected isAnagram(%q, %q) to be %v, but got %v", test.s, test.t, test.expected, result)
		}
	}
}
