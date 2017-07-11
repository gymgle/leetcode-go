package leetcode

import (
	"reflect"
	"testing"
)

func TestThreeSum_1(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}

	res := threeSum(nums)
	if reflect.DeepEqual(res, ret) {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}

func TestThreeSum_2(t *testing.T) {
	nums := []int{1, 1, -2}
	ret := [][]int{{-2, 1, 1}}

	res := threeSum(nums)
	if reflect.DeepEqual(res, ret) {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}
func TestThreeSum_3(t *testing.T) {
	nums := []int{0, 0, 0}
	ret := [][]int{{0, 0, 0}}

	res := threeSum(nums)
	if reflect.DeepEqual(res, ret) {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}

func TestThreeSum_4(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	ret := [][]int{{0, 0, 0}}

	res := threeSum(nums)
	if reflect.DeepEqual(res, ret) {
		t.Log("test passed")
	} else {
		t.Error("test failed")
	}
}
