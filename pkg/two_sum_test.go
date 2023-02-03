package pkg

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	ret := twoSum1([]int{1, 2, 3, 4}, 5)
	if ret != true {
		t.Fail()
	}
}

func TestTwoSum2(t *testing.T) {
	ret := twoSum2([]int{1, 7, 5, 4, 6, 3}, 10)
	if !reflect.DeepEqual(ret, [][]int{{4, 6}, {7, 3}}) {
		t.Fail()
	}
}
