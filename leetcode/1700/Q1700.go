package leetcode

func countStudents(students []int, sandwiches []int) int {
	var want0, want1 int
	for i := range students {
		if students[i] == 0 {
			want0++
		} else {
			want1++
		}
	}

	for i := range sandwiches {
		if sandwiches[i] == 0 {
			want0--
			if want0 < 0 {
				return want0 + want1 + 1
			}
		} else {
			want1--
			if want1 < 0 {
				return want0 + want1 + 1
			}
		}
	}
	return 0
}
