package bin_search

func Bin_search(A []int, key int) int{
	for left, right := 0, len(A)-1; left<=right; {
		mid := left + (right-left)/2
		if A[mid] == key {
			return mid
		} else if A[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/**
* 查找第一个等于给定值的元素
*/
func Bin_search_First_Equal(A []int, key int) int{
	for left, right := 0, len(A)-1; left<=right; {
		mid := left + (right-left)/2
		if A[mid] == key {
			if mid == 0 || A[mid-1] != key {
				return mid
			} else {
				right = mid - 1
			}
		} else if A[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}


/**
* 查找第一个大于等于给定值的元素
*/
func Bin_search_First_GT_Key(A []int, key int) int{
	for left, right := 0, len(A)-1; left<=right; {
		mid := left + (right-left)/2
		if A[mid] >= key {
			if mid == 0 || A[mid-1] < key {
				return mid
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}