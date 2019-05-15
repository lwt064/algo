package lis

var (
	a = []int{2,9,3,6,5,1,7}
	aLen = 7
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Reverse(x []int){
	nLen := len(x)
	for i, j := 0, nLen-1; i < j; {
		x[i], x[j] = x[j], x[i]
		i++
		j--
	}
}

func LIS() []int {
	maxLen := 0
	lisLen := make([]int, aLen)	// lisLen[i] = 以a[i]结尾的最长递增子序列长度
	for i, _ := range lisLen {
		lisLen[i] = 1
	}
	
	for i := 1; i < aLen; i++ {
		for j := i - 1; j >= 0; j-- {
			if a[j] <= a[i] {
				if lisLen[j] + 1 > lisLen[i] {
					lisLen[i] = lisLen[j] + 1
				}
				if lisLen[i] > maxLen {
					maxLen = lisLen[i]
				}
			}
		} 
	}

	lis := []int{}

	i := aLen - 1
	lastNum := 0
	for ; i >= 0; i-- {
		if lisLen[i] == maxLen {
			lastNum = a[i]
			lis = append(lis, lastNum)
			break
		}
	}

	maxLen--
	for ; i >= 0; i-- {
		if a[i] <= lastNum && lisLen[i] == maxLen {
			lastNum = a[i]
			lis = append(lis, lastNum)
			maxLen--
		}
	}
	Reverse(lis)
	return lis
}