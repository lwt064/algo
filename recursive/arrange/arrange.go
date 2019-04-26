package arrange

import "fmt"

func Arrange(A []int, k int, m int) {
	if k == m {
		fmt.Println(A)
	} else {
		for i := k; i < m; i++ {
			A[k], A[i] = A[i], A[k]
			Arrange(A, k+1, m)
			A[i], A[k] = A[k], A[i]
		}
	}
}
