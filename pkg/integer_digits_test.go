package pkg

import (
	"reflect"
	"testing"
)

func TestIntegerDigits1(t *testing.T) {
	ret := integerDigits1(87346)
	if !reflect.DeepEqual(ret, []int{8, 7, 3, 4, 6}) {
		t.Fail()
	}
}

func TestIntegerDigits2(t *testing.T) {
	ret := integerDigits2(87346)
	if !reflect.DeepEqual(ret, []int{8, 7, 3, 4, 6}) {
		t.Fail()
	}
}
