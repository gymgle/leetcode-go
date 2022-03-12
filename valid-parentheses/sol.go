package main

func isValid(s string) bool {
	if s == "" {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}
	m := map[byte]byte{'(': ')', '[': ']', '{': '}'}
	l := len(s)
	stack := make([]byte, 0)
	for i := 0; i < l; i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	println(isValid("{}"))
}
