Files
palindrome.go: Contains the implementation of the IsPalindrome function.
palindrome_test.go: Contains test cases for the IsPalindrome function to verify its correctness.
Function
IsPalindrome(s string) bool
The IsPalindrome function checks if the provided string s is a palindrome. A string is considered a palindrome if it reads the same forward and backward. The function compares characters from the beginning and end of the string moving towards the center.

Tests
The palindrome_test.go file includes the following test cases:

Palindrome Test: Checks if "radar" is correctly identified as a palindrome.
Non-Palindrome Test: Verifies that "hello" is not considered a palindrome.
Empty String Test: Confirms that an empty string is considered a palindrome.
Single Character Test: Ensures that a single character string is recognized as a palindrome.
