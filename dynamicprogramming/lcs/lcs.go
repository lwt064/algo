package lcs

var (
	a = []int{1,3,4,5,6,7,7,8}
	aLen = 8
	b = []int{3,5,7,4,8,6,7,8,2}
	bLen = 9
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func LCS() []interface{} {
	lcsLen := make([][]int, aLen+1)   // lcsLen[i][j] 长度为i的字符串与长度为j的字符串的编最长公共子序列长度
	for i, _ := range lcsLen {
		lcsLen[i] = make([]int, bLen+1)
		for j, _ := range lcsLen[i] {
			lcsLen[i][j] = 0
		}
	}
	
	for i := 1; i <= aLen; i++ {
		for j := 1; j <= bLen; j++ {
			if a[i-1] == b[j-1] {
				lcsLen[i][j] = lcsLen[i-1][j-1] + 1
			} else {
				l1 := lcsLen[i-1][j]
				l2 := lcsLen[i][j-1]
				lcsLen[i][j] = max(l1, l2)
			}
		}
	}

	// find one longest common sequence
	lcs := []interface{}{}
	for i := aLen; i > 0 ; {
		for j := bLen; j > 0 && i > 0; {
			if a[i-1] == b[j-1] {
				lcs = append(lcs, a[i-1])
				i--
				j--
			} else if lcsLen[i][j] == lcsLen[i-1][j] {
				i--
			} else if lcsLen[i][j] == lcsLen[i][j-1] {
				j--
			}
		}
	}
	return lcs
}