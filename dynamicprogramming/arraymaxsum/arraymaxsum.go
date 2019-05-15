package arraymaxsum

func DpArrayMaxSum() []int {
	b := []int{1, -2, 3, 10, -4, 7, 2, -5}

	state := make([]int, len(b))
	for i, _ := range b {
		state[i] = b[i]
	}

	sum := state[0]
	curfrom, curto := 0, 1
	bestfrom, bestto := curfrom, curto
	for i := 1; i < len(b); i++ {
		if state[i-1] > 0 {
			state[i] = state[i-1] + b[i]
			curto = i + 1
		} else {
			state[i] = b[i]
			curfrom = i
			curto = i + 1
		}

		if sum < state[i] {
			sum = state[i]
			bestfrom = curfrom
			bestto = curto
		}
	}

	return b[bestfrom:bestto]
}

func DpArrayMaxSum2() []int {
	b := []int{1, -2, 3, 10, -4, 7, 2, -5}

	sum, bestsum := b[0], b[0]
	curfrom, curto := 0, 1
	bestfrom, bestto := curfrom, curto
	for i := 1; i < len(b); i++ {
		if sum > 0 {
			sum += b[i]
			curto = i + 1
		} else {
			sum = b[i]
			curfrom = i
			curto = i + 1
		}

		if bestsum < sum {
			bestsum = sum
			bestfrom = curfrom
			bestto = curto
		}
	}

	return b[bestfrom:bestto]
}
