package editdist

var (
	a = "mitcmu"
	aLen = 6
	b = "mtacnu"
	bLen = 6
	MinDist = 99999999
)

func EditDist(i int, j int, dist int) {
	if i == aLen || j == bLen {
		if j < bLen {
			dist += bLen - j
		} else {
			dist += aLen - i
		}
		if MinDist > dist {
			MinDist = dist
		}
		return
	}

	if a[i] == b[j] {
		EditDist(i+1, j+1, dist)
	} else {
		EditDist(i+1, j, dist+1)
		EditDist(i, j+1, dist+1)
		EditDist(i+1, j+1, dist+1)
	}
}
