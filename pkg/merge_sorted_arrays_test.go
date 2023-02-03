package pkg

import (
	"reflect"
	"testing"
)

func TestMergeSortedArrays(t *testing.T) {
	l1 := []int{1, 6, 9}
	l2 := []int{2, 4, 5, 10}
	ret := mergeSortedArrays(l1, l2)
	if !reflect.DeepEqual([]int{1, 2, 4, 5, 6, 9, 10}, ret) {
		t.Fail()
	}
}
