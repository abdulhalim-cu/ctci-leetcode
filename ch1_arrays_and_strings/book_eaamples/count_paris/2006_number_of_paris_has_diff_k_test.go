package count_paris

import (
	"fmt"
	"testing"
)

func TestCountKDifference(t *testing.T) {
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 2, 1}, 1, 4},
		{[]int{1, 3}, 3, 0},
		{[]int{3, 2, 1, 5, 4}, 2, 3},
		{[]int{4, 4, 4, 4, 4}, 0, 10}, // All pairs have an absolute difference of 0
		{[]int{1, 2, 3, 4, 5}, 5, 0},  // No pairs with an absolute difference of 5
		{[]int{1, 2, 3, 4, 5}, 1, 4},  // Pairs: (1,2), (2,3), (3,4), (4,5)
		{[]int{1, 2, 3, 4, 5}, 0, 0},  // No pairs with an absolute difference of 0
		{[]int{1, 1, 1, 1, 1}, 0, 10}, // All pairs have an absolute difference of 0
		{[]int{1, 1, 1, 1, 1}, 1, 0},  // No pairs with an absolute difference of 1
		{[]int{1, 2, 3, 4, 5}, 10, 0}, // No pairs with an absolute difference of 10
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			got := countKDifference(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("Got %d, want %d", got, tc.want)
			}
		})
	}
}
