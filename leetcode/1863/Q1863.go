package leetcode

func subsetXORSum(nums []int) int {
	list := make([]int, 1, 1<<len(nums))
	for i := range nums {
		listLen := len(list)
		for j := 0; j < listLen; j++ {
			list = append(list, list[j])
			list[j] ^= nums[i]
		}
	}

	total := 0
	for _, n := range list {
		total += n
	}
	return total
}
