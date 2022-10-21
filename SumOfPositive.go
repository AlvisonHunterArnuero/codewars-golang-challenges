package kata

func PositiveSum(numbers []int) int {
	sum := 0
	for i := range numbers {
		if numbers[i] > 0 {
			sum += numbers[i]
		}
	}
	return sum
}
