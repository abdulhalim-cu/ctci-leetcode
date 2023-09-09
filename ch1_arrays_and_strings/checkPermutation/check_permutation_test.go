package checkPermutation

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	testCases := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"abc", "bca", true},
		{"abc", "def", false},
		{"hello", "loleh", true},
		{"good", "dog", false},
		{"world", "dlorwa", false},
		{"apple", "banana", false},
		{"", "", true},
		{"a", "a", true},
		{"abc", "abcd", false},
	}

	for _, tc := range testCases {
		t.Run(tc.s1+"_"+tc.s2, func(t *testing.T) {
			result := IsPermutation(tc.s1, tc.s2)
			if result != tc.expected {
				t.Errorf("For %s and %s, expected %t but got %t", tc.s1, tc.s2, tc.expected, result)
			}
		})
	}
}
