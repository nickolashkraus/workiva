package pkg

import "sort"

// Largest and Smallest Sum
//
// Given an array of 7 positive integers, find the largest value and the
// smallest value that can be calculated by summing 6 of the 7 elements of the
// array.
//
// NOTES
//   - Find min and max for O(n) time complexity.
func largestAndSmallestSum(arr []int) (int, int) {
	l, s := 0, 0
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for _, v := range arr[1:] {
		l = l + v
	}
	for _, v := range arr[:len(arr)-1] {
		s = s + v
	}
	return l, s
}
