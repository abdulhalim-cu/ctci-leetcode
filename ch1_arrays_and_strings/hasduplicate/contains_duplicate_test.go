package hasduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		// Test cases with duplicates
		{[]int{1, 2, 3, 4, 5, 1}, true},
		{[]int{9, 8, 7, 8, 6}, true},
		{[]int{1, 1, 1, 1, 1, 1}, true},

		// Test cases without duplicates
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 2, 3}, false},

		// Edge case with an empty slice
		{[]int{}, false},
	}

	for _, test := range tests {
		result := containsDuplicate(test.nums)
		if result != test.expected {
			t.Errorf("Expected containsDuplicate(%v) to be %v, but got %v", test.nums, test.expected, result)
		}
	}
}
