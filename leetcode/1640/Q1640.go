package leetcode

func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int]*[]int)
	for i := range pieces {
		m[pieces[i][0]] = &pieces[i]
	}
	for i := 0; i < len(arr); {
		head, ok := m[arr[i]]
		if !ok {
			return false
		}
		for j := 0; j < len(*head); j++ {
			if (*head)[j] != arr[i] {
				return false
			}
			i++
		}
	}
	return true
}
