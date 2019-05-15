package yanghui

func DpYanghui() int {
	n := 5
	Matrix := [5][5]int{
		{5, 8, 4, 1, 5},
		{7, 3, 6, 4},
		{2, 9, 9},
		{4, 7},
		{2},
	}

	state := make([][]int, n)
	for i, _ := range state {
		state[i] = make([]int, n)
		for j, _ := range state[i] {
			state[i][j] = 0
		}
	}

	state[0][0] = Matrix[0][0]
	for i := 1; i < n; i++ {
		state[i][0] = state[i-1][0] + Matrix[i][0]
	}
	for j := 1; j < n; j++ {
		state[0][j] = state[0][j-1] + Matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n-i; j++ {
			left := state[i][j-1]
			up := state[i-1][j]
			if left < up {
				state[i][j] = left + Matrix[i][j]
			} else {
				state[i][j] = up + Matrix[i][j]
			}
		}
	}

	minX := 99999999
	for i := 0; i < n; i++ {
		j := n - 1 - i
		if minX > state[i][j] {
			minX = state[i][j]
		}
	}
	return minX
}
