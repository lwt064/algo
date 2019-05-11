package bag01

func Resolve(w int, b []int) (int, []int) {
	max := 0
	selected := make([]int, len(b))
	bestselected := make([]int, len(b))
	for i, _ := range b {
		selected[i] = 0
		bestselected[i] = 0
	}

	Search(w, b, 0, &max, &selected, &bestselected)
	return max, bestselected
}

func Search(w int, b []int, cur int, max *int, selected *[]int, bestselected *[]int) {
	if cur == len(b) {
		sum := 0
		for _, x := range *selected {
			sum += x
		}
		if *max < sum {
			*max = sum
			for i, _ := range *selected {
				(*bestselected)[i] = (*selected)[i]
			}
		}
		return
	}
	(*selected)[cur] = 0
	Search(w, b, cur+1, max, selected, bestselected); // 不放当前物品

	sum := 0
	for _, x := range *selected {
		sum += x
	}
	if sum + b[cur] <= w {
		(*selected)[cur] = b[cur]
		Search(w, b, cur+1, max, selected, bestselected)
	}
}