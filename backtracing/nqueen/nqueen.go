package nqueen

import "math"
// import "fmt"

func Resolve(n int) int {
	pos := make([]int, n)
	path := 0
	Place(0, n, &pos, &path)
	return path
}

func Place(row int, maxrow int, pos *[]int, path *int) {
	if row == maxrow {
		// fmt.Println(*pos)
		*path++
		return
	}
	maxcol := maxrow
	for j := 0; j < maxcol; j++ {   // 每行有 maxcol 种放法
		if IsOk(row, j, *pos) {
			(*pos)[row] = j
			Place(row+1, maxrow, pos, path)
			(*pos)[row] = 0			// 可以不要，因为下一次调用会覆盖
		}
	}
}

func IsOk(row, col int, pos []int) bool {
	for i := row - 1; i >= 0; i-- {
		if pos[i] == col {
			return false
		}
		diff := math.Abs(float64(row-i)) - math.Abs(float64(col-pos[i]))
		if math.Abs(diff) < 1 {
			return false
		}
	}
	return true
}
