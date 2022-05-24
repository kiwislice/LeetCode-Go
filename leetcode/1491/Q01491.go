package leetcode

func average(salary []int) float64 {
	min := 1000001
	max := 0
	var sum float64
	for _, n := range salary {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += float64(n)
	}
	sum -= float64(min)
	sum -= float64(max)
	return sum / float64(len(salary)-2)
}
