package _2

import (
	"testing"
)

func TestGCDOfStrings(t *testing.T) {
	// Test case 1: Common divisor "ABC"
	str1 := "ABCABC"
	str2 := "ABC"
	result := gcdOfStrings(str1, str2)
	expected := "ABC"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 2: Common divisor "AB"
	str1 = "ABABAB"
	str2 = "ABAB"
	result = gcdOfStrings(str1, str2)
	expected = "AB"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 3: No common divisor
	str1 = "LEET"
	str2 = "CODE"
	result = gcdOfStrings(str1, str2)
	expected = ""
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 4: Empty string, should return an empty string
	str1 = "X"
	str2 = "XYZ"
	result = gcdOfStrings(str1, str2)
	expected = ""
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}

	// Test case 5: Equal strings, common divisor is the string itself
	str1 = "HELLO"
	str2 = "HELLO"
	result = gcdOfStrings(str1, str2)
	expected = "HELLO"
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}
