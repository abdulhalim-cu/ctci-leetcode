package hashTable

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	// Test cases where ransomNote can be constructed from magazine
	t.Run("Test1", func(t *testing.T) {
		ransomNote := "a"
		magazine := "b"
		expected := false
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test2", func(t *testing.T) {
		ransomNote := "aa"
		magazine := "ab"
		expected := false
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test3", func(t *testing.T) {
		ransomNote := "aa"
		magazine := "aab"
		expected := true
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	// Additional test cases
	t.Run("Test4", func(t *testing.T) {
		ransomNote := "abc"
		magazine := "defabcxyz"
		expected := true
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test5", func(t *testing.T) {
		ransomNote := "apple"
		magazine := "applaespp"
		expected := true
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

	t.Run("Test6", func(t *testing.T) {
		ransomNote := "hello"
		magazine := "world"
		expected := false
		result := canConstruct(ransomNote, magazine)
		if result != expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
