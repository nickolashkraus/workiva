package pkg

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	s := "racecar"
	ret := isPalindrome(s)
	if !ret {
		t.Fail()
	}
	s = "leetcode"
	ret = isPalindrome(s)
	if ret {
		t.Fail()
	}
}
