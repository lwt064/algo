package coin

func DpCoin() int {
	b := []int{1, 3, 5}
	v := 9

	state := make([]int, v+1)
	for i := 0; i <= v; i++ {
		state[i] = 0
	}
	for i := 0; i <= v; i++ {
		for _, curVal := range b {
			if i < curVal {
				break
			}
			if state[i] == 0 || state[i-curVal]+1 < state[i] {
				state[i] = state[i-curVal] + 1
			}
		}
	}
	return state[v]
}
