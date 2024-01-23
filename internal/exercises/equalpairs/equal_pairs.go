package equalpairs

func EqualPairs(args ...interface{}) interface{} {
	var grid = args[0].([][]int)

	return equalPairs(grid)
}

func equalPairs(grid [][]int) int {
	n := len(grid)

	rows := make([][]int, n)
	for r := 0; r < n; r++ {
		rows[r] = grid[r]
	}

	cols := make([][]int, n)
	for r := 0; r < n; r++ {
		list := make([]int, n)
		for c := 0; c < n; c++ {
			list[c] = grid[c][r]
		}
		cols[r] = list
	}

	counter := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if equalSlices(rows[i], cols[j]) {
				counter++
			}
		}
	}

	return counter
}

func equalSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
