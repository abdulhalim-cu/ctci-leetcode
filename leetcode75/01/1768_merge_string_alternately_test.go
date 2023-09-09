package _1

import (
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	// Test case 1: Basic test case with equal length strings
	result := mergeAlternately("abc", "pqr")
	expected := "apbqcr"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 2: First string shorter than the second
	result = mergeAlternately("ab", "pqrs")
	expected = "apbqrs"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 3: Second string shorter than the first
	result = mergeAlternately("abcd", "pq")
	expected = "apbqcd"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 4: First string is empty
	result = mergeAlternately("", "xyz")
	expected = "xyz"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 5: Second string is empty
	result = mergeAlternately("123", "")
	expected = "123"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 6: Both strings are empty
	result = mergeAlternately("", "")
	expected = ""
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
