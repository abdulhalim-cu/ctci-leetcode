package permutations

import (
	"reflect"
	"testing"
)

func TestUniqPermutations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aab", []string{"aab", "aba", "baa"}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			result := uniqPermutations(testCase.input)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("For input %s, expected %v, but got %v", testCase.input, testCase.expected, result)
			}
		})
	}
}
