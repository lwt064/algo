package arrange

import "fmt"

func Arrange(data []int, k int)  {
	if k == len(data) {
		fmt.Println(data)
		return
	}
	for i := k; i < len(data); i++ {
		data[i], data[k] = data[k], data[i]
		Arrange(data, k+1)
		data[i], data[k] = data[k], data[i]
	}
}