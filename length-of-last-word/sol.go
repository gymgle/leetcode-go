package main

func lengthOfLastWord(s string) int {
	count := 0
	lenght := len(s)
	start := lenght - 1
	for start >= 0 {
		if s[start] == ' ' {
			start--
			continue
		}
		break
	}
	for start >= 0 {
		if s[start] != ' ' {
			count++
			continue
		}
		break
	}
	return count
}
