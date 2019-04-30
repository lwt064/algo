package cellsplit

// 1个细胞生命周期三小时，每小时分裂一次

var (
	x1 = []int{1, 2, 4, 7}
	x2 = []int{1, 2, 4, 8}
)

// 假设第三小时分裂前死亡
// f(n) = 2 * f(n-1) - f(n-3)
// f(0) = 1
// f(1) = 2
// f(2) = 4
// f(3) = 7
// f(4) = 12
func Split1(n int) int {
	if n < len(x1) {
		return x1[n]
	}
	return 2*Split1(n-1) - Split1(n-3)
}

// 假设第三小时分裂后再死亡
// f(n) = 2 * f(n-1) - f(n-4)
// f(0) = 1
// f(1) = 2
// f(2) = 4
// f(3) = 8
// f(4) = 15
func Split2(n int) int {
	if n < len(x2) {
		return x2[n]
	}
	return 2*Split2(n-1) - Split2(n-4)
}
