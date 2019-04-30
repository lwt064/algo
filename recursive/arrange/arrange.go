package arrange

import "fmt"

/**
* m个数全排队问题，相当于固定第一个数后，m-1个数的全排列问题
* @param A  数组
* @param m 数组长度
* @param k 当前开始处理的位置
**/
func Arrange(A []int, m int, k int) {
	if k == m {
		fmt.Println(A)
	} else {
		for i := k; i < m; i++ {
			A[k], A[i] = A[i], A[k]
			Arrange(A, m, k+1)
			A[i], A[k] = A[k], A[i]
		}
	}
}
