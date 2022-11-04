package pkg

// This question is literally 125. Valid Palindrome:
//   - https://leetcode.com/problems/valid-palindrome
//
// NOTES
//   - Use two pointers.
//
// This question was so easy I was almost embarrassed I received it for a
// Senior Software Engineer role. The interviewers also suggested I implement
// the same function recursively.
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func isPalindromeRecursive(s string) bool {
	if s[0] != s[len(s)-1] {
		return false
	} else if len(s) <= 2 {
		return true
	}
	return isPalindromeRecursive(s[1 : len(s)-1])
}
