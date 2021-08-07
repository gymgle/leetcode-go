package nqueens

import (
	"strings"
)

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	res := [][]string{}
	colums := map[int]bool{}
	diag1 := map[int]bool{}
	diag2 := map[int]bool{}
	helper(0, bd, &res, n, colums, diag1, diag2)
	return res
}

func helper(r int, bd [][]string, res *[][]string, n int, colums, diag1, diag2 map[int]bool) {
	if r == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, tmp)
		return
	}

	for c := 0; c < n; c++ {
		if !colums[c] && !diag1[r-c] && !diag2[r+c] {
			bd[r][c] = "Q"
			colums[c] = true
			diag1[r-c] = true
			diag2[r+c] = true
			helper(r+1, bd, res, n, colums, diag1, diag2)
			bd[r][c] = "."
			colums[c] = false
			diag1[r-c] = false
			diag2[r+c] = false
		}
	}
}
