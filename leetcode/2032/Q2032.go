package leetcode

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	counts := [100 + 1]int{}
	for _, n := range nums1 {
		counts[n] |= 1
	}
	for _, n := range nums2 {
		counts[n] |= 2
	}
	for _, n := range nums3 {
		counts[n] |= 4
	}
	rtn := make([]int, 0, 10)
	for i, n := range counts {
		if n == 3 || n > 4 {
			rtn = append(rtn, i)
		}
	}
	return rtn
}
