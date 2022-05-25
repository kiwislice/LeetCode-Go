package leetcode

func cellsInRange(s string) []string {
	a, z := s[0], s[3]
	i, j := s[1], s[4]
	w, h := z-a+1, j-i+1

	rtn := make([]string, w*h)
	for c := a; c <= z; c++ {
		for r := i; r <= j; r++ {
			rtn[(c-a)*h+r-i] = string(c) + string(r)
		}
	}
	return rtn
}
