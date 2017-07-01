package leetcode

import "testing"

func Test_removeDuplicates_1(t *testing.T) {
	array := []int{1, 2}
	if removeDuplicates(array) == 2 {
		t.Log("1st test passed")
	} else {
		t.Error("1st test failed")
	}
}

func Test_removeDuplicates_2(t *testing.T) {
	array := []int{1, 1, 1, 2, 2, 3}
	if removeDuplicates(array) == 5 {
		t.Log("2nd test passed")
	} else {
		t.Error("2nd test failed")
	}
}
