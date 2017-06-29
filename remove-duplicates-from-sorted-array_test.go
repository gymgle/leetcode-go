package leetcode

import "testing"

func Test_removeDuplicates(t *testing.T) {
	array1 := []int{}
	array2 := []int{1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 4}

	if removeDuplicates(array1) != 0 {
		t.Error("1st test faild")
	} else {
		t.Log("1st test passed")
	}

	if removeDuplicates(array2) != 4 {
		t.Error("2nd test failed")
	} else {
		t.Log("2nd test passed")
	}
}
