package leetcode

func getLucky(s string, k int) int {
	rtn := 0
	for _, c := range s {
		n := int(c - 'a' + 1)
		rtn += n/10 + n%10
	}
	for i := 1; i < k; i++ {
		rtn = sum(rtn)
	}
	return rtn
}

func sum(n int) int {
	rtn := 0
	for n > 0 {
		rtn += n % 10
		n /= 10
	}
	return rtn
}
