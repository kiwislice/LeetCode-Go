package leetcode

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	open := make([]bool, n)

	// var dfs func(from int)
	// dfs = func(from int) {
	// 	open[from] = true
	// 	for _, to := range rooms[from] {
	// 		if !open[to] {
	// 			dfs(to)
	// 		}
	// 	}
	// }

	dfs(rooms, 0, open)
	for _, opened := range open {
		if !opened {
			return false
		}
	}
	return true
}

func dfs(rooms [][]int, from int, open []bool) {
	open[from] = true
	for _, to := range rooms[from] {
		if !open[to] {
			dfs(rooms, to, open)
		}
	}
}
