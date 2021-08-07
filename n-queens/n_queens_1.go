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
	helper(0, bd, &res, n)
	return res
}

func helper(r int, bd [][]string, res *[][]string, n int) {
	if r == n {
		tmp := make([]string, n)
		for i := 0; i < n; i++ {
			tmp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, tmp)
		return
	}

	for c := 0; c < n; c++ {
		if isValid(r, c, n, bd) {
			bd[r][c] = "Q"
			helper(r+1, bd, res, n)
			bd[r][c] = "."
		}
	}
}

func isValid(r, c, n int, bd [][]string) bool {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (c == j || r+c == i+j || r-c == i-j) {
				return false
			}
		}
	}
	return true
}
