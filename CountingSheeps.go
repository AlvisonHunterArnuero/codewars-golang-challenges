package kata

func CountSheeps(numbers []bool) int {
	cnt := 0
	for _, value := range numbers {
		if value {
			cnt++
		}
	}
	return cnt
}
