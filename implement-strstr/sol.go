package main

import "strings"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}
	return strings.Index(haystack, needle)
}
