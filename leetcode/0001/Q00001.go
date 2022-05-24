package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, n := range nums {
		v, ok := m[n]
		if ok {
			return []int{v, i}
		} else {
			m[target-n] = i
		}
	}
	return []int{}
}
