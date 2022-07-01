package leetcode

import "container/list"

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	group := make([]int, n)
	groupCount := 1
	queue := list.New()

	for i := 0; i < n; i++ {
		if group[i] > 0 {
			continue
		}

		group[i] = groupCount
		queue.PushBack(i)
		for queue.Len() > 0 {
			p := queue.Remove(queue.Front()).(int)
			for j := 0; j < n; j++ {
				if isConnected[p][j] == 1 && group[j] == 0 {
					group[j] = groupCount
					queue.PushBack(j)
				}
			}
		}
		groupCount++
	}
	return groupCount - 1
}
