package dpeditdist

var (
	a = "mitcmu"
	aLen = len(a)
	b = "mtacnu"
	bLen = len(b)
	MinDist = 99999999
)

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func DpEditDist() int {
	dist := make([][]int, aLen+1)   // dist[i][j] 长度为i的字符串与长度为j的字符串的编辑距离
	for i, _ := range dist {
		dist[i] = make([]int, bLen+1)
		for j, _ := range dist[i] {
			dist[i][j] = 0
		}
	}
	for i := 1; i <= aLen; i++ {
		dist[i][0] = i
	}
	for j := 1; j <= bLen; j++ {
		dist[0][j] = j
	}

	for i := 1; i <= aLen; i++ {
		for j := 1; j <= bLen; j++ {
			if a[i-1] == b[j-1] {
				dist[i][j] = dist[i-1][j-1]
			} else {
				d1 := dist[i-1][j] + 1
				d2 := dist[i][j-1] + 1
				d3 := dist[i-1][j-1] + 1
				dist[i][j] = min(min(d1, d2), d3)
			}
		}
	}

	return dist[aLen][bLen]
}