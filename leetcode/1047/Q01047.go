package leetcode

func removeDuplicates(s string) string {
	stack := make([]rune, len(s))
	i := 0
	for _, c := range s {
		if i > 0 && stack[i-1] == c {
			i--
		} else {
			stack[i] = c
			i++
		}
	}
	return string(stack[:i])
}
