package leetcode

func eventualSafeNodes(graph [][]int) []int {
	UNKNOW, UNSAFE, SAFE := 0, 1, 2
	n := len(graph)
	status := make([]int, n)

	var isSafe func(from int) bool
	isSafe = func(from int) bool {
		if status[from] != UNKNOW {
			return status[from] == SAFE
		}

		status[from] = UNSAFE
		for _, to := range graph[from] {
			if !isSafe(to) {
				return false
			}
		}
		status[from] = SAFE
		return true
	}

	safeNodes := make([]int, 0, n)
	for p := range graph {
		if isSafe(p) {
			safeNodes = append(safeNodes, p)
		}
	}
	return safeNodes
}

func eventualSafeNodes_TimeOut(graph [][]int) []int {
	UNSAFE, SAFE := 1, 2
	n := len(graph)
	status := make([]int, n)

	path := make(map[int]bool)
	var dfsMarkCircle func(from int)
	dfsMarkCircle = func(from int) {
		if status[from] == UNSAFE {
			return
		}
		path[from] = true
		for _, to := range graph[from] {
			if status[to] == UNSAFE || path[to] {
				for p := range path {
					if path[p] {
						status[p] = UNSAFE
					}
				}
				return
			} else if status[to] == SAFE {
				continue
			} else {
				dfsMarkCircle(to)
			}
		}

		if status[from] != UNSAFE {
			status[from] = SAFE
		}
		path[from] = false
	}

	for p := range graph {
		dfsMarkCircle(p)
	}

	safeNodes := make([]int, 0, n)
	for i := range status {
		if status[i] != UNSAFE {
			safeNodes = append(safeNodes, i)
		}
	}
	return safeNodes
}
