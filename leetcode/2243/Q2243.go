package leetcode

func digitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}

	ar := []rune(s)
	for i := range ar {
		ar[i] -= '0'
	}
	for len(ar) > k {
		ar = round(ar, k)
	}
	for i := range ar {
		ar[i] += '0'
	}

	return string(ar)
}

func round(s []rune, k int) []rune {
	lenS := len(s)
	groupCount := (lenS + k - 1) / k
	i := 0

	for g := 0; g < groupCount; g++ {
		from := g * k
		end := from + k
		if end > lenS {
			end = lenS
		}

		var subsum rune
		for i := from; i < end; i++ {
			subsum += s[i]
		}

		from = i
		if subsum == 0 {
			s[i] = 0
			i++
		} else {
			for subsum > 0 {
				s[i] = subsum % 10
				subsum /= 10
				i++
			}
		}

		end = i - 1
		reverse(s, from, end)
	}
	return s[:i]
}

func reverse(s []rune, from, to int) {
	for from < to {
		temp := s[from]
		s[from] = s[to]
		s[to] = temp
		from++
		to--
	}
}
