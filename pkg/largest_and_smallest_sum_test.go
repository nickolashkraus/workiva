package pkg

import (
	"testing"
)

func TestLargestAndSmallestSum(t *testing.T) {
	arr := []int{3, 2, 5, 12, 8, 7, 10}
	l, s := largestAndSmallestSum(arr)
	if l != 45 {
		t.Fail()
	}
	if s != 35 {
		t.Fail()
	}
}
