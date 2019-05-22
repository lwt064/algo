package array

import "fmt"

const MAXNUM = 100

func GetSeqOfS(s int) {
	A := make([]int, MAXNUM+1)
	for i := 0; i <= MAXNUM; i++ {
		A[i] = i
	}

	small := 1
	big := 2
	sum := 3
	if s < sum {
		fmt.Println("#")
		return
	}

	found := false
	for small <= (s+1)/2 {
		if sum == s {
			fmt.Println(A[small : big+1])
			found = true
		}
		if sum <= s {
			big = big + 1
			sum += big
		} else {
			sum -= small
			small = small + 1
		}
	}

	if !found {
		fmt.Println("#")
	}
}

func PrintMatrix() {
	m, n := 8, 6
	A := make([][]int, m)

	num := 0
	for i := 0; i < m; i++ {
		A[i] = make([]int, n)
		for j := 0; j < n; j++ {
			num += 1
			A[i][j] = num
		}
	}

	for _, x := range A {
		fmt.Println(x)
	}

	k := 1 // k为圈数
	row, col := 0, 0

	cnt := 0
	for {
		cnt = 0
		for ; col < n-k; col++ {
			cnt += 1
			fmt.Printf("%d ", A[row][col])
		}
		if cnt == 0 {
			break
		}

		cnt = 0
		for ; row < m-k; row++ {
			cnt = cnt + 1
			fmt.Printf("%d ", A[row][col])
		}
		if cnt == 0 {
			break
		}

		cnt = 0
		for ; col >= k; col-- {
			cnt = cnt + 1
			fmt.Printf("%d ", A[row][col])
		}
		if cnt == 0 {
			break
		}

		cnt = 0
		for ; row >= k; row-- {
			cnt = cnt + 1
			fmt.Printf("%d ", A[row][col])
		}
		row++
		col++
		if cnt == 0 {
			break
		}

		k = k + 1
	}

}
