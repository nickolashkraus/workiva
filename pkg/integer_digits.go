package pkg

import (
	"container/list"
	"strconv"
)

// Integer Digits
//
// Create a function that takes a int and returns a list of its digits.
//
// EXAMPLE:
//
//	Input: 87346
//	Output: [8, 7, 3, 4, 6]
func integerDigits1(i int) []int {
	ret := []int{}
	// Convert int to str.
	s := strconv.Itoa(i)
	for _, c := range s {
		i, _ := strconv.Atoi(string(c))
		ret = append(ret, i)
	}
	return ret
}

// Alternative solution that does not convert the integer to a string.
//
// GOTCHA: The array is in the reverse order.
//
// FOLLOW-UP:
//   - What data structure is most efficient for insertions and deletions?
func integerDigits2(i int) []int {
	l := list.New()
	for i > 0 {
		l.PushFront(i % 10)
		i = i / 10
	}
	ret := make([]int, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		if val, ok := e.Value.(int); ok {
			ret = append(ret, val)
		}
	}
	return ret
}
