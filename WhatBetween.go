package kata

func Between(a, b int) []int {
	arr := make([]int, b-a+1)
	for i := range arr {
		arr[i] = a + i
	}
	return arr
}
