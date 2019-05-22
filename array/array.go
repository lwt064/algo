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
