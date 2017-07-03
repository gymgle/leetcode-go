package leetcode

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := [][]int{
		[]int{0, 1},
		[]int{1, 2, 3},
		[]int{2, 7, 11, 15},
	}

	targets := []int{
		1,
		4,
		22,
	}

	results := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 3},
	}

	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := twoSum(tests[i], targets[i]); ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Error("test failed")
		} else {
			t.Log("test passed")
		}
	}
}
