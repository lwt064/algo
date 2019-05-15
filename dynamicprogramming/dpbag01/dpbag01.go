package dpbag01

var (
	b []int // weight，整数数组
	v []int // value
	k int   // k个物品
	w int   // 最大重量，整数
)

func InitBag() {
	b = []int{5, 4, 7, 2, 6}
	v = []int{12, 3, 10, 3, 6}
	k = 5
	w = 15
}

func Bag01() int {
	state := make([][]int, k+1)
	for i, _ := range state {
		state[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			state[i][j] = 0
		}
	}

	for i := 1; i <= k; i++ { // 前i个物品放入背包
		curW := b[i-1]
		curV := v[i-1]
		for j := 0; j <= w; j++ {
			if j < curW {
				state[i][j] = state[i-1][j]
			} else {
				if state[i-1][j] > state[i-1][j-curW]+curV {
					state[i][j] = state[i-1][j]
				} else {
					state[i][j] = state[i-1][j-curW] + curV
				}
			}
		}
	}

	maxv := 0
	for _, x := range state[k] {
		if maxv < x {
			maxv = x
		}
	}
	return maxv
}

func Bag01_2() int {
	state := make([]int, w+1)
	for j, _ := range state {
		state[j] = 0
	}

	for i := 1; i <= k; i++ { // 前i个物品放入背包
		curW := b[i-1]
		curV := v[i-1]
		for j := w; j >= 0; j-- { // state[j]依赖state[j-curW]，所以需要从右往左计算，不然state[j-curW]会被破坏
			if j < curW {
				// nothing
			} else {
				if state[j] > state[j-curW]+curV {
					// nothing
				} else {
					state[j] = state[j-curW] + curV
				}
			}
		}
	}

	maxv := 0
	for _, x := range state {
		if maxv < x {
			maxv = x
		}
	}
	return maxv
}
