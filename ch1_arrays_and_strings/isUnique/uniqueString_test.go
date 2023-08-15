package isunique

import "testing"

func TestCheckHasUniqueChar(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		// Test cases with duplicates
		{"abcdefghijklmnopqrstuvwxyz", true},
		{"helloworld", false},
		{"áéíóúáé", false},
		{"ßüëçã", true},
		{"ååå", false},
		{"®©¢¥", true},
		{"AΔ!@", true},
		{"±±±±", false},
		{"ÆØÅæøå", true},
		{"", true},
	}

	for _, test := range tests {
		result := CheckHasUniqueChar(test.s)
		if result != test.expected {
			t.Errorf("Expected containsDuplicate(%v) to be %v, but got %v", test.s, test.expected, result)
		}
	}
}

func TestCheckHasUniqueCharUsingBitwiseOperation(t *testing.T) {
	// Test case 1: All characters are unique
	input1 := "abcdefghijklmnopqrstuvwxyz"
	result1 := CheckHasUniqueCharUsingBitwiseOperation(input1)
	if !result1 {
		t.Errorf("Test case 1 failed: Expected true, got false")
	}

	// Test case 2: Repeated character 'e'
	input2 := "helloworld"
	result2 := CheckHasUniqueCharUsingBitwiseOperation(input2)
	if result2 {
		t.Errorf("Test case 2 failed: Expected false, got true")
	}

	// Test case 3: String with only one character
	input3 := "a"
	result3 := CheckHasUniqueCharUsingBitwiseOperation(input3)
	if !result3 {
		t.Errorf("Test case 3 failed: Expected true, got false")
	}

	// Test case 4: Empty string
	input4 := ""
	result4 := CheckHasUniqueCharUsingBitwiseOperation(input4)
	if !result4 {
		t.Errorf("Test case 4 failed: Expected true, got false")
	}

	// Test case 5: String with repeated characters
	input5 := "abbcc"
	result5 := CheckHasUniqueCharUsingBitwiseOperation(input5)
	if result5 {
		t.Errorf("Test case 5 failed: Expected false, got true")
	}

	// Test case 6: String with repeated characters
	input6 := "abcdabcd"
	result6 := CheckHasUniqueCharUsingBitwiseOperation(input6)
	if result6 {
		t.Errorf("Test case 6 failed: Expected false, got true")
	}
}
