package leetcode

import (
	"sort"
)

func frequencySort(nums []int) []int {
	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}
	less := func(i, j int) bool {
		if counts[nums[i]] == counts[nums[j]] {
			return nums[i] > nums[j]
		}
		return counts[nums[i]] < counts[nums[j]]
	}
	sort.Slice(nums, less)
	return nums
}
