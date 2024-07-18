// palindrome_test.go
package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	// Test case: Palindrome
	if !IsPalindrome("radar") {
		t.Error(`IsPalindrome("radar") = false; want true`)
	}

	// Test case: Not a palindrome
	if IsPalindrome("hello") {
		t.Error(`IsPalindrome("hello") = true; want false`)
	}

	// Test case: Empty string
	if !IsPalindrome("") {
		t.Error(`IsPalindrome("") = false; want true`)
	}

	// Test case: Single character
	if !IsPalindrome("a") {
		t.Error(`IsPalindrome("a") = false; want true`)
	}
}
