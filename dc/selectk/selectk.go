package selectk

import "sort"

func Partition(S *[][]int, mid int) {
	i := 0
	j := len(*S) - 1
	for i <= j {
		for i < len(*S) && (*S)[i][len((*S)[i])/2] <= mid {
			i++
		}
		for j >= 0 && (*S)[j][len((*S)[j])/2] > mid {
			j--
		}
		if i >= j {
			break
		}
		(*S)[i], (*S)[j] = (*S)[j], (*S)[i]
		i++
		j--
	}
}

// A左闭右开
func SelectK(S []int, k int) int {
	nLen := len(S)
	if nLen < 10 && k < 10 {
		tmp := S
		sort.Ints(tmp)
		return tmp[nLen-k]
	}

	S2 := [][]int{}
	for i := 0; i < nLen; i += 5 {
		end := i + 5
		if end > nLen {
			end = nLen
		}

		tmp := S[i:end]
		sort.Ints(tmp)
		S2 = append(S2, tmp)
	}

	midA := make([]int, len(S2))
	for i, _ := range S2 {
		midA[i] = S2[i][len(S2[i])/2]
	}

	mid := SelectK(midA, len(midA)/2)

	Partition(&S2, mid)

	// device S2 to A B C D
	A, B, C, D := []int{}, []int{}, []int{}, []int{}
	for i := 0; i < len(S2); i++ {
		if S2[i][len(S2[i])/2] <= mid {
			for j := 0; j <= len(S2[i])/2; j++ {
				A = append(A, S2[i][j])
			}
			for j := len(S2[i])/2 + 1; j < len(S2[i]); j++ {
				B = append(B, S2[i][j])
			}
		} else {
			for j := 0; j < len(S2[i])/2; j++ {
				C = append(C, S2[i][j])
			}
			for j := len(S2[i]) / 2; j < len(S2[i]); j++ {
				D = append(D, S2[i][j])
			}
		}
	}
	A = A[0 : len(A)-1] // remove mid

	// A <= m, D > m
	S3 := A
	S4 := D
	for _, x := range B {
		if x <= mid {
			S3 = append(S3, x)
		} else {
			S4 = append(S4, x)
		}
	}
	for _, x := range C {
		if x <= mid {
			S3 = append(S3, x)
		} else {
			S4 = append(S4, x)
		}
	}

	if k == len(S4)+1 {
		return mid
	} else if k <= len(S4) {
		return SelectK(S4, k)
	} else {
		return SelectK(S3, k-len(S4)-1)
	}
}
