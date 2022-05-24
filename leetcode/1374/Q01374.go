package leetcode

func generateTheString(n int) string {
	ar := make([]rune, n)
	if n&0x1 == 0 {
		ar[0] = 'a'
		for i := 1; i < n; i++ {
			ar[i] = 'b'
		}
	} else {
		for i := 0; i < n; i++ {
			ar[i] = 'a'
		}
	}
	return string(ar)
}
