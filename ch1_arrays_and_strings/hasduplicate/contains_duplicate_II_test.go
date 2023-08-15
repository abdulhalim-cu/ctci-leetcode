package hasduplicate

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	// Test case 1: Distinct elements, but within k distance
	input1 := []int{1, 2, 3, 1}
	k1 := 3
	result1 := ContainsNearbyDuplicate(input1, k1)
	if !result1 {
		t.Errorf("Test case 1 failed: Expected true, got false")
	}

	// Test case 2: Duplicate elements within k distance
	input2 := []int{1, 0, 1, 1}
	k2 := 1
	result2 := ContainsNearbyDuplicate(input2, k2)
	if !result2 {
		t.Errorf("Test case 2 failed: Expected true, got false")
	}

	// Test case 3: No duplicate elements within k distance
	input3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	result3 := ContainsNearbyDuplicate(input3, k3)
	if result3 {
		t.Errorf("Test case 3 failed: Expected false, got true")
	}

	// Test case 4: Empty array
	input4 := []int{}
	k4 := 0
	result4 := ContainsNearbyDuplicate(input4, k4)
	if result4 {
		t.Errorf("Test case 4 failed: Expected false, got true")
	}

	// Test case 5: Large array with no duplicates
	input5 := []int{5, 10, 15, 20, 25}
	k5 := 2
	result5 := ContainsNearbyDuplicate(input5, k5)
	if result5 {
		t.Errorf("Test case 5 failed: Expected false, got true")
	}
}
