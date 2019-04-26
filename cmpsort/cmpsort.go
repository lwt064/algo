package cmpsort

func MergeSort(A []int) []int{
	if len(A) < 2 {
		return A
	}
	mid := len(A)/2
	left := MergeSort(A[:mid])
	right := MergeSort(A[mid:])
	return merge(left, right)
}

func merge(left []int, right []int) []int{
	i, j := 0, 0
	result := make([]int, 0)
	for ; i < len(left) && j < len(right); {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i = i+1
		} else {
			result = append(result, right[j])
			j = j+1
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}
	for ; j < len(right); j++ {
		result = append(result, right[j])
	}
	return result
}

func HeapSort(A []int) {
	if len(A) < 2 {
		return
	}
	s := len(A)/2
	for ; s >=0 ; s-- {
		siftDown(A, s, len(A))
	}

	for i := len(A)-1; i>=0; i-- {
		A[0], A[i] = A[i], A[0]
		siftDown(A, 0, i)
	}
}

func siftDown(A []int, s int, length int) {
	dad := s
	son := 2*s + 1
	for ; son < length; {
		if son+1 < length && A[son] < A[son+1] {
			son = son + 1
		}
		if A[dad] < A[son] {
			A[son], A[dad] = A[dad], A[son]
			dad = son
			son = 2*son + 1
		} else {
			break
		}
	}
}

func QuicSort(A []int) {
	if len(A) < 2 {
		return
	}
	mid := partition(A)
	QuicSort(A[:mid])
	QuicSort(A[mid+1:])
}


func partition(A []int) int {
	if len(A) < 2{
		return 0
	}
	e := len(A) - 1
	pivot := A[e]
	i, j := 0, 0
	for ; j < e ; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i = i+1
		}
	}
	A[i], A[e] = A[e], A[i]
	return i
}