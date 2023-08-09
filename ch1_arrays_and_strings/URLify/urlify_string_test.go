package URLify

import (
	"fmt"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Mr John Smith", "Mr%20John%20Smith"},
		{"  Hello World  ", "Hello%20World"},
		{"  a b c d  ", "a%20b%20c%20d"},
		{"", ""},
		{"A", "A"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %q", tc.input), func(t *testing.T) {
			result := ReplaceSpace(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %q, Got: %q", tc.expected, result)
			}
		})
	}
}
